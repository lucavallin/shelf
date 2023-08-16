---
showToc: true
TocOpen: false
hidemeta: false
comments: true
ShowCanonicalLink: false
disableShare: false
searchHidden: true
hideSummary: false
ShowReadingTime: true
ShowBreadCrumbs: true
ShowPostNavLinks: true
# ShowWordCount: true
ShowRssButtonInSectionTermList: true
UseHugoToc: true
editPost:
    URL: "https://github.com/lucavallin/lucavallin.github.io/tree/main/content"
    Text: "Suggest Changes"
    appendFilePath: true
# CanonicalLinkText
# disableAnchoredHeadings
# hideFooter
# ShowCodeCopyButtons
# ShareButtons
# robotsNoIndex

author: "Luca Cavallin"
date: "2021-03-04"
title: "How to optimize PHP performance on Google Cloud Run"
description: "I recently had to optimize the performance of a PHP-based API on Cloud Run. After a performance test, we discovered that the API became very slow when we put some serious load on it (with response times exceeding 10 seconds). In this post you’ll learn what changes I made to get that down to a stable 100ms."
canonicalURL: "https://binx.io/blog/2021/03/04/optimizing-php-performance-google-cloudrun/"
tags: ["google cloud", "cloud run", "php", "nginx", "docker"]
draft: false
cover:
    image: "/luca.png"
    alt: "Luca Cavallin | Software Engineer"
    caption: "Luca Cavallin | Software Engineer" # display caption under cover
    relative: false # when using page bundles set this to true
    hidden: false # only hide on current single page
# weight: 1
# aliases: ["/first"]
---

The API uses PHP 7.4, Laravel 8.0 and MySQL on Cloud SQL (the managed database on Google Cloud). The API needs to handle at least 10,000 concurrent users. The container image we deploy to Cloud Run has nginx and PHP-FPM.

This is what I did to improve the response times. (Don’t worry if not everything makes sense here, I’ll explain everything).
- Matching the number of PHP-FPM workers to the maximum concurrency setting on Cloud Run.
- Configuring OPcache (it compiles and caches PHP scripts)
- Improving composer auto-loading settings
- Laravel-specific optimizations including caching routes, views, events and using API resources

### Matching Concurrency and Workers
The application uses nginx and PHP-FPM, which is a process manager for PHP. PHP is single threaded, which means that one process can handle one (exactly one) request at the same time. PHP-FPM keeps a pool of PHP workers (a worker is a process) ready to serve requests and adds more if the demand increases.
It’s a good practice to limit the maximum size of the PHP-FPM worker pool, to make sure your resource usage (CPU and memory) is predictable.

To configure PHP-FPM for maximum performance, I first set the process manager type for PHP-FPM to static, so that the specified number of workers are running at all times and waiting to handle requests. I did this by copying a custom configuration file to the application’s container and configuring the environment so that these options will be picked up by PHP-FPM (you must copy the configuration where it is expected, in my case, into `/usr/local/etc/php-fpm.d/`). The settings I needed are:
```
pm = static
pm.max_children = 10
```
However, if you set a limit, and more requests come to a server than the pool can handle, requests start to queue, which increases the response time of those requests:

<img src="https://binx.io/wp-content/uploads/2021/03/nginx-php-fpm-900x592.png" alt="Nginx and php-fpm request model" width="900" height="592" class="size-medium wp-image-2970" />

### Limiting Concurrent Requests on Cloud Run
To avoid request queuing in nginx, you’ll need to limit the number of requests Cloud Run sends to your container at the same time.

Cloud Run uses request-based autoscaling. This means it limits the amount of concurrent requests it sends to a container, and adds more containers if all containers are at their limit. You can change that limit with the [concurrency setting](https://cloud.google.com/run/docs/about-concurrency "concurrency setting"). I set it to 10, which I determined is the maximum number of concurrent requests a container with 1GB of memory and 1vCPU can take for with this application.

![Cloud Run concurrent requests with Nginx and php-fpm](https://binx.io/wp-content/uploads/2021/03/conc-reqs-900x848.png)

You really want to make sure Cloud Run's concurrency setting matches the maximum number of PHP-FPM workers! For example, if Cloud Run sends 100 concurrent requests to a container before adding more containers, and you configured your PHP-FPM to start only 10 workers, you will see a lot of requests queuing.

If these tweaks aren’t enough to reach the desired performance, check the Cloud Run metrics to see what the actual utilization percentages are. You might have to change the amount of memory and vCPUs available to the container. The downside of this optimization is that more containers will be running due to lower concurrency, resulting in higher costs. I also noticed temporary delays when new instances are starting up, but this normalizes over time.


### Configuring OPCache
OPCache is a default PHP extension that caches the compiled scripts in memory, improving response times dramatically. I enabled and tweaked OPCache settings by adding the extension’s options to a custom `php.ini` file (in my case, I put it in the `/usr/local/etc/php/conf.d/` directory). The following is a generic configuration that you can easily reuse, and you can refer to the [documentation](https://www.php.net/manual/en/opcache.configuration.php "documentation") for the details about every option.

```
opcache.enable=1
opcache.memory_consumption=128
opcache.interned_strings_buffer=64
opcache.max_accelerated_files=32531
opcache.validate_timestamps=0
opcache.save_comments=1
opcache.fast_shutdown=0
```

### Optimizing Composer
Composer is a dependency manager for PHP. It lets you specify the libraries your app needs, and downloads them for you to a directory. It also generates an autoload configuration file, which maps import paths to files.
If you pass the `--optimize-autoloader` flag to composer, it will generate this file only once, and doesn’t dynamically update it if you add new code. While that is convenient in development (your changes show up immediately), in production it can make your code really slow.

You can [optimize](https://getcomposer.org/doc/articles/autoloader-optimization.md#autoloader-optimization "optimize") Composer’s autoloader passing the `--optimize-autoloader` flag like this:
```
composer install --optimize-autoloader --no-dev
```
### Laravel-specific optimizations
The application I optimized is built with Laravel, which provides a number of tools that can help improving the performance of the API. Here's what I did on top of the other tweaks to get the response times below 100ms.

- I have leveraged Laravel’s built-in caching features during builds to reduce start-up times. There are no downsides to these tweaks, except that you won’t be able to use closure-defined routes (they can’t be cached). You can cache views, events and routes with these commands:

 ```
php artisan view:cache
php artisan event:cache
php artisan route:cache
```
Avoid running `php artisan config:cache` since Laravel ignores environment variables if you cache the configuration.

- Using Laravel [API Resources](https://laravel.com/docs/8.x/eloquent-resources "API Resources") further improves the response times of your application. This has proven to be much faster than having the framework automatically convert single objects and collections to JSON.

### Summary
In this blog, I shared with you what I learned from optimizing the performance of a PHP-based API on Cloud Run. All of the tweaks together helped cut response times to one tenth of the original result, and I think the most impact was made by matching concurrency and PHP-FPM workers (if you’re in a hurry, do only this). Watching the application metrics has been fundamental throughout the performance testing phase, just as inspecting Cloud Run logs after each change.

If your application still shows poor performance after these changes, there are other tweaks you can make to improve response times, which I haven't discussed here.
- Increase PHP memory limits if needed
- Check MySQL for slow queries (often due to missing indexes)
- Cache the responses with a CDN
- Migrate to PHP 8.0 (up to 3x [faster](https://www.php.net/releases/8.0/en.php "faster"))
