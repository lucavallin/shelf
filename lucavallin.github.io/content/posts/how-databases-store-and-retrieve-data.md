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
date: "2023-07-14"
title: "How Databases Store and Retrieve Data with B-Trees"
description: >
    This is what I learnt by reading the book "Database Internals: A Deep Dive Into How Distributed Data Systems Work" by Alex Petrov (O'Reilly Media). Besides storing and retrieving data, there's a lot more a database does: it manages concurrency, ensures durability, parses and executes queries, handles transactions, and so on. This post focuses on the storage engine, which is responsible for storing and retrieving data efficiently.
canonicalURL: ""
tags: ["database", "b-tree", "storage engine", "mysql"]
draft: false
cover:
    image: "/images/posts/how-databases-store-and-retrieve-data/cover.png"
    alt: "How Databases Store and Retrieve Data with B-Trees"
    caption: "How Databases Store and Retrieve Data with B-Trees"
    relative: false # when using page bundles set this to true
    hidden: false # only hide on current single page
# weight: 1
# aliases: ["/first"]
---

**Databases** are the backbone of modern applications, enabling efficient data storage, retrieval, and manipulation. Behind the scenes, database storage engines utilise sophisticated data structures, one of which is the **B-tree**. I recently read "_Database Internals: A Deep Dive Into How Distributed Data Systems Work_" to learn more about the topic, and in this blog post I am going to summarise how database storage engines leverage B-trees to handle data and the challenges they face along the way.

## The B-Tree: an overview

Think of a B-tree as a hierarchical structure resembling an upside-down tree. Its branches (nodes) extend downwards, with leaves representing the actual data. Each node in the B-tree contains keys that separate the data stored within it. These keys facilitate efficient search and retrieval operations, even with large amounts of data.

### Insertion and Deletion Mechanics

One of the primary challenges in managing databases is efficient insertion and deletion of data. B-trees excel in addressing this challenge. To insert new data, a database's storage engine follows rules to maintain the tree's balance and order. It starts at the root node, comparing the new key with existing keys until it finds the appropriate leaf node. The new data seamlessly fits into the correct position while preserving the sorted order.

Similarly, when deleting data, the storage engine carefully removes the target key and adjusts the remaining keys to **maintain balance and integrity**. This dynamic nature of B-trees allows for efficient insertion and deletion operations, ensuring that the database remains performant despite frequent modifications.

#### A practical example with MySQL

MySQL, a popular relational database management system, utilizes B-trees as its primary data structure for storing and retrieving data efficiently. Let's explore a practical example of how MySQL leverages B-trees in the context of data storage and retrieval.

Consider a table named "Customers" with the following structure:

```sql
CREATE TABLE Customers (
    id INT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    address VARCHAR(200)
);
```

When data is inserted into the "Customers" table, MySQL uses B-trees to organize and store the data efficiently on disk. Each row in the table represents a data entry, containing information such as the customer's ID, name, email, and address.
Internally, MySQL maintains a B-tree index, often referred to as the **primary index or clustered index**, on the primary key column ("id" in this case). This B-tree structure maps the primary key values to the corresponding data rows, ensuring fast and direct access to individual records.

### Efficient Search and Range Queries

In addition to insertion and deletion, B-trees excel at supporting search queries. Retrieving specific data from a database involves traversing the B-tree based on the provided search key. Starting from the root node, the engine compares the search key with the keys in each node and determines the appropriate direction. By following this path, the engine efficiently narrows down the search space until it reaches the desired leaf node, enabling rapid data retrieval.

Furthermore, B-trees effortlessly handle range queries, which involve retrieving data within a specified range of keys. Since the keys are sorted within each node, the storage engine can quickly identify the starting and ending points of the desired range. It then traverses the tree to gather the necessary data, making range queries a fast and efficient process.

#### Another practical example with MySQL

MySQL leverages B-trees to efficiently retrieve data based on specified search conditions. For example, if we want to retrieve customer records with a specific email address, we can use the following query:

```sql
SELECT * FROM Customers WHERE email = 'example@example.com';
```

MySQL utilizes the B-tree index associated with the primary key to quickly locate the corresponding data entry that matches the specified email value. This fast data retrieval is possible because the B-tree's hierarchical structure allows for efficient traversal to the desired leaf node, where the relevant data resides.

B-trees in MySQL also facilitate efficient range queries, enabling the retrieval of data within a specified range of values. For instance, if we want to retrieve customers with IDs between 100 and 200, we can use the following query:

```sql
SELECT * FROM Customers WHERE id BETWEEN 100 AND 200;
```

MySQL leverages the B-tree index on the primary key to efficiently identify the leaf nodes that contain the desired range of IDs. By traversing the B-tree index, MySQL can swiftly retrieve the corresponding data entries, optimizing the range query execution.

As data is modified in the "Customers" table through insertions, updates, or deletions, MySQL's B-tree index dynamically adjusts to reflect these changes. The B-tree is **rebalanced and reorganized** as necessary to maintain its balanced structure and optimize performance. This ensures that the B-tree remains efficient for subsequent data storage and retrieval operations.

![Tree](/images/posts/how-databases-store-and-retrieve-data/tree.png)

## The Disk

Databases generally store data on disk rather than in memory. Let's take a closer look at how B-trees are stored and read, considering the challenges of disk I/O.

### Disk Storage Structure

When a B-tree is stored on disk, it is typically divided into fixed-size pages or _blocks_. Each page corresponds to a node in the B-tree. The size of a page is determined by the storage system and is usually a multiple of the disk's sector size.

The B-tree's root node is stored in a fixed location on disk, often referred to as the root page. Each internal node of the B-tree consists of a set of keys and pointers to child pages. The child pages, in turn, contain more keys and pointers, forming the hierarchical structure of the tree. Finally, the leaf pages hold the actual data entries.

### Reading from Disk

When a query or operation requires accessing a B-tree, the storage engine must retrieve the necessary pages from disk. This process involves disk I/O, which can be a potential bottleneck due to the relatively slower speed of disk access compared to memory.

To minimize the number of disk reads and optimize performance, storage engines employ various techniques:

- **Caching**: A cache, such as a buffer pool, is used to hold frequently accessed pages in memory. By keeping frequently accessed pages in the cache, subsequent reads can be performed directly from memory, reducing the need for disk I/O.
- **Page Pre-fetching**: Storage engines often employ predictive algorithms to anticipate which pages are likely to be accessed in the near future. By pre-fetching these pages into the cache, the engine can reduce the latency associated with fetching pages on-demand.
- **I/O Optimization**: Storage engines employ strategies to optimize disk I/O, such as reading multiple pages at once in sequential or asynchronous fashion, reducing seek times. Additionally, techniques like write-ahead logging (WAL) are used to group multiple modifications into a single disk write, improving overall efficiency.
By combining these techniques, storage engines strive to minimize the impact of disk I/O and enhance the performance of reading B-trees from disk.

#### The usual practical example with MySQL

In MySQL, tables are stored on disk using a storage engine. MySQL supports multiple storage engines, such as InnoDB, MyISAM, Memory, etc., each with its own method of storing data. The choice of storage engine determines how tables are stored on disk. Let's take a look at InnoDB as an example:

- InnoDB uses a file-per-table approach, where each table is stored in its own separate file. The files are typically stored in the MySQL data directory.
- Each InnoDB table is divided into _pages_ (usually 16KB in size), and the pages are stored in a tablespace. The tablespace consists of one or more data files, known as InnoDB data files (.ibd files).
- The tablespace can be shared among multiple tables, and it manages the storage, caching, and concurrency control of the data.
- InnoDB uses a clustered index structure, where the data is physically stored based on the primary key or the first unique index defined on the table.
- Additional indexes defined on the table are stored separately from the data in a separate B+tree structure.

Indexes in MySQL use pointers to locate the corresponding data rows. The exact mechanism of how indexes point to data rows depends on the storage engine being used:

- In InnoDB, the primary key is also known as the clustered index. The clustered index determines the physical order of rows in the table.
- InnoDB uses a B+tree structure to organize the clustered index. The leaf nodes of the B+tree contain the actual data rows of the table, and the non-leaf nodes contain index key values and pointers to child nodes.
- When a query is executed using a condition that matches an index, the InnoDB storage engine uses the B+tree structure to efficiently traverse the index and locate the corresponding data rows.
- The leaf nodes of the clustered index contain a special value called the "row ID" that identifies the physical location of the data row within the tablespace.

##### Index Maintenance

When a record is inserted or removed in MySQL, the offsets or pointers in the indexes typically need to be updated to reflect the changes in the physical location of the data rows. The process of updating these offsets is known as index maintenance. Let's consider two scenarios:

###### Record Insertion

- When a new record is inserted into a table, the storage engine determines the appropriate location to place the record based on the index structure.
- If the table has one or more indexes, the storage engine needs to update the index entries to include the newly inserted record.
- In most cases, the index structure needs to be modified to accommodate the new record. This may involve adding a new leaf node in the B+tree structure or modifying existing nodes to make room for the new entry.
- Additionally, if the inserted record affects the order of the index key values, the affected nodes in the index structure may need to be adjusted to maintain the sorted order.
- The **offsets** or **pointers** in the index structure are updated to point to the newly inserted record's physical location.

###### Record Deletion

- When a record is deleted from a table, the storage engine needs to remove the corresponding entry from the index structure.
- The storage engine locates the index entry associated with the deleted record and removes it from the index structure.
If the deletion affects the order of the index key values, the affected nodes in the index structure may need to be adjusted to maintain the sorted order.
- After removing the index entry, the index **offsets** or **pointers** are updated to reflect the changes in the physical location of the remaining data rows.

## Challenges in B-Tree Management

While B-trees offer numerous benefits, they also pose challenges. Here are a few notable obstacles faced by database storage engines when using B-trees:

- **Balancing Act**: Maintaining the balance of a B-tree is crucial for efficient data access. As the database grows and shrinks, rebalancing the tree becomes necessary to prevent it from becoming skewed. Ensuring that nodes have a roughly equal number of keys requires careful splitting and merging operations, which can introduce overhead.
- **Disk I/O Bottlenecks**: B-trees are often stored on disk, not in memory, which introduces the challenge of minimizing disk I/O. Efficiently reading and writing data from disk involves optimising structure serialization, page accesses and caching mechanisms to minimize costly disk seeks, enabling faster data retrieval.
- **Concurrent Access**: In multi-user database environments, concurrent access by multiple threads or processes can introduce concurrency issues. B-trees need to handle concurrent read and write operations safely to maintain data consistency and integrity.

![Databass](/images/posts/how-databases-store-and-retrieve-data/bass.png)

### Conclusion

B-trees play a vital role in database storage engines, offering efficient data storage, retrieval, and manipulation. They tackle the challenges of insertion, deletion, search, and range queries effectively, making them a popular choice for managing large volumes of data. Storing and reading B-trees from disk involves dividing the tree into **fixed-size pages** and employing caching, pre-fetching, and I/O optimization techniques to minimize disk I/O and enhance performance. These strategies allow the storage engine to efficiently access the necessary pages while mitigating the inherent latency associated with disk access. B-trees, with their balanced structure and disk storage optimizations, continue to be a reliable choice for managing large-scale databases, providing efficient data retrieval and manipulation while gracefully handling the challenges of disk I/O.
