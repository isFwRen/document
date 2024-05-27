### 基础篇

#### 一、Select语句执行顺序

**sql的执行顺序**

在去写好sql和优化sql查询之前，知道sql的执行顺序尤为的重要（所以这里要多读几遍）

**from ->on ->join ->where ->group by ->having ->select ->distinct ->order by ->limit**
1、from 对查询指定的表计算笛卡尔积
2、on 按照 join_condition 过滤数据
3、join 添加关联外部表数据
4、where 按照where_condition过滤数据
5、group by 进行分组操作
6、having 按照having_condition过滤数据
7、select 选择指定的列
8、distinct 指定列去重
9、order by 按照order_by_condition排序
10、limit 取出指定记录量

![image-20240402135612977](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240402135612977.png)

-  **select \*** : 返回所有记录
-  **limit N** : 返回 N 条记录
-  **offset M** : 跳过 M 条记录, 默认 M=0, 单独使用似乎不起作用
-  **limit N,M** : 相当于 **limit M offset N** , 从第 N 条记录开始, 返回 M 条记录

#### 二、MySQL创建数据库

登录MySQL服务后，使用`CREATE`关键字创建数据库，**关键字建议大写，可读性比较高**，语法如下：

```
CREATE DATABASE 数据库名;
```

建数据库的基本语法如下：

```
CREATE DATABASE [IF NOT EXISTS]database_name
	[CHARACTET SET charset_name]
	[COLLATE collation_name];
```

如果希望在创建数据库时指定一些选项，可以使用 `CREATE DATABASE`语句的其他参数，例如可以指定字符集和排序规则：

```
CREATE DATABASE mydatabase
	CHARACTER SET utf8mb4
	COLLATE utf8mb4_general_ci;
```

如果数据库已经存在，执行CREATE DATABASE将导致错误。为了避免这种情况，可以在CREATE DATABASE语句中添加ID NOT EXISTS字句：

```
CREATE DATABASE IF NOT EXISTS mydatabase
	CHARACTER SET utf8mb4
	COLLATE utf8mb4_general_ci;
```

这条命令的作用：

- 如果数据库不存在则创建，存在则不创建
- 创建mydatabase数据库，并设定编码集为utf8和排序规则

#### 三、DROP删除数据库

DORP命令格式：

```
DROP DATABASE <database_name>;	--直接删除数据库，不检查是否存在
或
DROP DATABASE [IF EXISTS] <database_name>;
```

参数说明：

- IF EXIDSTS是一个可选子句，表示如果数据库存在才执行删除操作，避免因为数据库不存在而引发错误
- database_name是你要删除的数据库名称。

示例：

```
--直接删除数据库，不检查是否存在
DEOP DATABASE mydatabase;

--删除数据库，如果存在的话
DROP DATABASE IF EXISTS mydatabase;
```

**注意**：在执行删除数据库操作前，请确保你确实想要删除数据库及其所有数据，因为该操作是不可逆的。为了避免误操作，通常建议执行删除之前备份数据库。

#### 四、查看和选择数据库

登录MySQL服务后，使用`show databases;`命令查看连接上服务的全部数据库

选择要操作的数据库使用的命令：`use mydata_name;`

#### 五、MySQL 数据类型

MySQL 中定义数据字段的类型对你数据库的优化是非常重要的。

MySQL 支持多种类型，大致可以分为三类：数值、日期/时间和字符串(字符)类型。

****

**数值类型**

MySQL 支持所有标准 SQL 数值数据类型。

这些类型包括严格数值数据类型(INTEGER、SMALLINT、DECIMAL 和 NUMERIC)，以及近似数值数据类型(FLOAT、REAL 和 DOUBLE PRECISION)。

关键字INT是INTEGER的同义词，关键字DEC是DECIMAL的同义词。

BIT数据类型保存位字段值，并且支持 MyISAM、MEMORY、InnoDB 和 BDB表。

作为 SQL 标准的扩展，MySQL 也支持整数类型 TINYINT、MEDIUMINT 和 BIGINT。下面的表显示了需要的每个整数类型的存储和范围。

| 类型         | 大小                                     | 范围（有符号）                                               | 范围（无符号）                                               | 用途            |
| :----------- | :--------------------------------------- | :----------------------------------------------------------- | :----------------------------------------------------------- | :-------------- |
| TINYINT      | 1 Bytes                                  | (-128，127)                                                  | (0，255)                                                     | 小整数值        |
| SMALLINT     | 2 Bytes                                  | (-32 768，32 767)                                            | (0，65 535)                                                  | 大整数值        |
| MEDIUMINT    | 3 Bytes                                  | (-8 388 608，8 388 607)                                      | (0，16 777 215)                                              | 大整数值        |
| INT或INTEGER | 4 Bytes                                  | (-2 147 483 648，2 147 483 647)                              | (0，4 294 967 295)                                           | 大整数值        |
| BIGINT       | 8 Bytes                                  | (-9,223,372,036,854,775,808，9 223 372 036 854 775 807)      | (0，18 446 744 073 709 551 615)                              | 极大整数值      |
| FLOAT        | 4 Bytes                                  | (-3.402 823 466 E+38，-1.175 494 351 E-38)，0，(1.175 494 351 E-38，3.402 823 466 351 E+38) | 0，(1.175 494 351 E-38，3.402 823 466 E+38)                  | 单精度 浮点数值 |
| DOUBLE       | 8 Bytes                                  | (-1.797 693 134 862 315 7 E+308，-2.225 073 858 507 201 4 E-308)，0，(2.225 073 858 507 201 4 E-308，1.797 693 134 862 315 7 E+308) | 0，(2.225 073 858 507 201 4 E-308，1.797 693 134 862 315 7 E+308) | 双精度 浮点数值 |
| DECIMAL      | 对DECIMAL(M,D) ，如果M>D，为M+2否则为D+2 | 依赖于M和D的值                                               | 依赖于M和D的值                                               | 小数值          |

------

**日期和时间类型**

表示时间值的日期和时间类型为DATETIME、DATE、TIMESTAMP、TIME和YEAR。

每个时间类型有一个有效值范围和一个"零"值，当指定不合法的MySQL不能表示的值时使用"零"值。

TIMESTAMP类型有专有的自动更新特性，将在后面描述。

| 类型      | 大小 ( bytes) | 范围                                                         | 格式                | 用途                     |
| :-------- | :------------ | :----------------------------------------------------------- | :------------------ | :----------------------- |
| DATE      | 3             | 1000-01-01/9999-12-31                                        | YYYY-MM-DD          | 日期值                   |
| TIME      | 3             | '-838:59:59'/'838:59:59'                                     | HH:MM:SS            | 时间值或持续时间         |
| YEAR      | 1             | 1901/2155                                                    | YYYY                | 年份值                   |
| DATETIME  | 8             | '1000-01-01 00:00:00' 到 '9999-12-31 23:59:59'               | YYYY-MM-DD hh:mm:ss | 混合日期和时间值         |
| TIMESTAMP | 4             | '1970-01-01 00:00:01' UTC 到 '2038-01-19 03:14:07' UTC结束时间是第 **2147483647** 秒，北京时间 **2038-1-19 11:14:07**，格林尼治时间 2038年1月19日 凌晨 03:14:07 | YYYY-MM-DD hh:mm:ss | 混合日期和时间值，时间戳 |

------

**字符串类型**

字符串类型指CHAR、VARCHAR、BINARY、VARBINARY、BLOB、TEXT、ENUM和SET。该节描述了这些类型如何工作以及如何在查询中使用这些类型。

| 类型       | 大小                  | 用途                            |
| :--------- | :-------------------- | :------------------------------ |
| CHAR       | 0-255 bytes           | 定长字符串                      |
| VARCHAR    | 0-65535 bytes         | 变长字符串                      |
| TINYBLOB   | 0-255 bytes           | 不超过 255 个字符的二进制字符串 |
| TINYTEXT   | 0-255 bytes           | 短文本字符串                    |
| BLOB       | 0-65 535 bytes        | 二进制形式的长文本数据          |
| TEXT       | 0-65 535 bytes        | 长文本数据                      |
| MEDIUMBLOB | 0-16 777 215 bytes    | 二进制形式的中等长度文本数据    |
| MEDIUMTEXT | 0-16 777 215 bytes    | 中等长度文本数据                |
| LONGBLOB   | 0-4 294 967 295 bytes | 二进制形式的极大文本数据        |
| LONGTEXT   | 0-4 294 967 295 bytes | 极大文本数据                    |

**注意**：char(n) 和 varchar(n) 中括号中 n 代表字符的个数，并不代表字节个数，比如 CHAR(30) 就可以存储 30 个字符。

CHAR 和 VARCHAR 类型类似，但它们保存和检索的方式不同。它们的最大长度和是否尾部空格被保留等方面也不同。在存储或检索过程中不进行大小写转换。

BINARY 和 VARBINARY 类似于 CHAR 和 VARCHAR，不同的是它们包含二进制字符串而不要非二进制字符串。也就是说，它们包含字节字符串而不是字符字符串。这说明它们没有字符集，并且排序和比较基于列值字节的数值值。

BLOB 是一个二进制大对象，可以容纳可变数量的数据。有 4 种 BLOB 类型：TINYBLOB、BLOB、MEDIUMBLOB 和 LONGBLOB。它们区别在于可容纳存储范围不同。

有 4 种 TEXT 类型：TINYTEXT、TEXT、MEDIUMTEXT 和 LONGTEXT。对应的这 4 种 BLOB 类型，可存储的最大长度不同，可根据实际情况选择。

------

**枚举与集合类型（Enumeration and Set Types）**

- **ENUM**: 枚举类型，用于存储单一值，可以选择一个预定义的集合。
- **SET**: 集合类型，用于存储多个值，可以选择多个预定义的集合。

------

**空间数据类型（Spatial Data Types）**

GEOMETRY, POINT, LINESTRING, POLYGON, MULTIPOINT, MULTILINESTRING, MULTIPOLYGON, GEOMETRYCOLLECTION: 用于存储空间数据（地理信息、几何图形等）。

****

**拓展：**

MySQL 5.0 以上的版本：

1、一个汉字占多少长度与编码有关：

**UTF－8**：一个汉字＝3个字节

**GBK**：一个汉字＝2个字节

2、varchar(n) 表示 n 个字符，无论汉字和英文，Mysql 都能存入 n 个字符，仅是实际字节长度有所区别

3、MySQL 检查长度，可用 SQL 语言来查看：

```
select LENGTH(fieldname) from tablename
```

#### 六、MySQL创建数据表

创建MySQL数据表需要以下信息：

- 表明
- 表字段名
- 定义每个表字段的数据类型

**语法**

下面是创建MySQL数据表的SQL通用语法：

```
CREATE TABLE table_name (
	colum1 datatype,
	colum2 datatype,
	...
);
```

**参数说明：**

- table_name是你要创建的表的名称；
- colum1，colum2，...是表中的列名；
- datatype是每个列的数据类型。

以下是一个具体的实例，创建一个用户表users：

```
CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(50) NOT NULL,
	email VARCHAR(100) NOT NULL,
	birthdate DATE,
	is_active BOOLEAN DEFAULT TRUE
);
```

实例解析：

- id：用户id。整型类型，自增长，作为主键。
- username：用户名，边长字符串，不允许为空。
- email：用户邮箱，边长字符串，不允许为空。
- birthdate：用户的生日，日期类型。
- is_active：用户是否已经激活，布尔类型，默认值为true。

以上是一个简单的实例，用到了一些常用的数据类型包括INT,VARCHAR,DATE,BOOLEAN，你可以根据实际需要选择不同的数据类型。AUTO_INCREMENT关键字用于创建一个自主增长的列，PRIMARY KEY用于定于主键。

如果希望在创建表时指定数据引擎，字符集和排序规则等，可以使用CHARCTER SET 和 COLLATE子句：

```
CREATE TABLE IF NOT EXISTS `runoob_tb1` (
	`runoob_id`	INT UNSIGEND AUTO_INCREMENT,
	`runoob_title` VARCHAR(100) NOT NULL,
	`runoob_author` VARCHAR(40) NOT NULL,
	`submission_date` DATE,
	PRIMARY KEY (`runoob_id`)
)ENGINE=InnoDB DEFAULT CHAERSET=utf8;
```

实例解析：

- 如果你不想字段为空可以设置字段的属性为NOT NILL，在操作数据库时如果输入改字段的数据为空，就会报错。
- **AUTO_INCREMENT**定义列为自增的属性，一般用于主键，数值会自动加1.
- **PRIMARY KEY**关键字用于定于列为主键。可以定于多列定义主键，列间以逗号`，`分割。
- **ENGINE**设置存储引擎，**CHARSET**设置编码。

**注意：**

创建 MySql 的表时，表名和字段名外面的符号  **`**  不是单引号，而是英文输入法状态下的反单引号，也就是键盘左上角 **esc** 按键下面的那一个 **~** 按键，坑惨了。

反引号是为了区分 MySql 关键字与普通字符而引入的符号，一般的，表名与字段名都使用反引号。

#### 七、MySQL删除数据表

MySQL中删除数据表是非常容易操作的，但是你在进行删除表操作时要非常小心，因为执行删除命令后所有数据都会消失。

**语法**

以下为删除MySQL数据表的通用语法：

```
DROP TABLE table_name;	--直接删除表，不检查是否存在
或
DROP TABLE [IF EXISTS] table_name;
```

**参数说明：**

- table_name是要删除的表的名称。
- IF EXISTS是一个可选的子句，表示如果存在才执行删除操作，避免因为表不存在而引发错误。

```
--删除表，如果存在的话
DROP TABLE IF EXISTS mytable;

--直接删除表，不检查是否存在
DROP TABLE mytable;
```

#### 八、MySQL插入数据

MySQL表中使用**INSERT INTO**语句来插入数据。

你可以通过mysql>命令提示窗口向数据表中插入数据，或者通过PHP脚本来插入数据。

**语法**

以下未向MySQL数据表插入数据通用的**INSERT INTO**语法：

```
INSERT INT table_name (colum1,colum2,colum3,...)
VALUES (value1,value2,value3,...);
```

**参数说明：**

- table_name是你要插入数据的表的名称。
- colum1,colum2,colum3,...是表中的列名（字段名）。
- value1,value2,value3,...是要插入的具体数值。

如果数据是字符型，必须使用单引号 `'`或者双引号`"`,如：'value1',"value1"。

一个简单的实例，插入了一行数据到名为users的表中：

```
INSERT INTO users (username,email,birthdate,is_active)
VALUES ('test','test@runoob.com','1990-01-01',true);
```

- username:用户名，字符串类型。
- email:邮箱地址，字符串类型。
- birthdate：用户生日，日期类型。
- is_active：是否已激活，布尔类型。

如果你想要插入所有列的数据，可以省略列名：

```
INSERT INTO users
VALUES (NULL,'test','test@runoob.com','1999-01-02',true)
```

这里的**NULL**是用于自增长列的占位符，表示系统将为id列生成一个唯一的值。

如果你要插入多行数据，可以在VALUES子句中指定多组数值：

```
INSERT INTO users (username,email,birthdate,is_active)
VALUES
	('test1','test1@runoob.com','1999-01-01',true),
	('test1','test1@runoob.com','1999-01-01',true),
	('test1','test1@runoob.com','1999-01-01',true);
```

以上代码将在suers表中插入三行数据。

#### 九、MySQL查询数据

MySQL数据库使用**SELECT**语句来查询数据。

**语法**

以下为在MySQL数据库中拆线呢数据通用的**SELECT**语法：

```
SELECT colum1,colum2,...
FROM table_name
[WHERE condition]
[ORDER BY column_name [ASC|DESC]]
[LIMIT number];
```

**参数说明：**

- column1,column2,...是你想要选择的列的名称，如果使用`*`表示选择所有列。
- table_name是你要从中查询数据的表的名称。
- WHERE condition 是一个可选子句，用于指定过滤条件，只返回复符合条件的行。
- ORDER BY column_name [ASC | DESC ] 是一个i可选子句，用于指定结果集的排序顺序，，默认是升序（ASC）。
- LIMIT number是一个可选的子句，用于限制返回的行数。

MySQL SELECT语句简单的应用实例：

```
--选择所有列的所有行
SELECT * FROM users;

--选择特定列的所有行
SELECT username,email FROM users;

--添加WHERE 子句，按照满足条件的行
SELECT * FROM users WHERE is_active = TRUE;

--添加ORDER BY 子句，按照某列的升序排序
SELECT * FROM users ORDER BY birthdate;

--添加ORDER BY 子句，按照某些的降序排序
SELECT * FROM users ORDER BY birthdate DESC;

--添加LIMIT 子句，限制返回的行数
SELECT * FROM users LIMIT 10;
```

SELECT语句可以是灵活的，我们可以根据实际需要组合和使用这些子句，比如同时使用WHERE和ORDER BY 子句，或者使用LIMIT控制返回的行数。

在WHERE子句中，你可以使用各种条件运算符（如=，<,>,<=,>=,!=），逻辑运算符（如AND，OR，NOT），以及通配符（如%，_）等。

以下是一些进阶的SELECT语句实例：

```
--使用AND运算符和通配符
SELECT * FROM users WHERE username LIKE 'J%' AND is_active = TRUE;

--使用OR运算符
SELECT * FROM users WHERE is_active = TRUE OR birthdate < '1999-01-01';

--使用IN子句
SELECT * FROM users WHERE birthdate IN ('1990-01-01','1992-01-01','1993-01-05');
```

**MySQL limit应用的一些例子：**

语法格式：

```
SELECT * FROM table LIMIT [offset,] rows OFFSET offset
```

解析：LIMIT 子句可以被用于强制SELECT语句返回指定的记录数。LIMITE接收**一个或者两个参数**。参数必须是一个整数常量。如果给定两个参数，第一个参数指定第一个返回记录行的偏移量，第二个参数指定返回记录行的最大数目。初始记录行的偏移量是0（而不是1）：为了鱼PostgreSQL兼容，MySQL也支持句法：LIMIT#OFFSET#.

```
SELECT * FROM table LIMIT 5,10;  //检索记录行6-15

//如果只给定一个参数，它表示返回最大的记录行数码：
SELECT * FROOM table LIMIT 5; //检索前5个记录行

//换句话说，LIMIT n 等价于 LMITE 0,n。
```

**MySQL的分页查询语句的性能分析**

MySQL分页sql语句，如果和MSSQL的TOP语法相比，那么MySQL的LIMIT语法显得优雅了许多。使用它来分页再自然不过的事情了。

2.1最基本的分页方式：

```
SELECT ... FROM ... WHERE ... ORDER BY ... LIMIT ...
```

在中小数据量的情况下，这样的SQL足够用了，唯一需要注意的问题是确保使用了索引。

举例来说，如果实际SQL类似下面语句，那么在category_id，id两列上建立复合索引比较好。

```
SELECT * FROM articles WHERE category_id = 123 ORDER BY id LIMIT 50,10
```

#### 十、MySQL WHERE子句

我们知道从MySQL表中使用**SELECT**语句来读取数据。

如需有条件地从表中选取数据，可将**WHERE**子句添加到**SELECT**语句中。

**WHERE**子句用于在MySQL中过滤查询结果，只返回满足特定条件的行。

**语法**

以下是SQL SELECT语句使用WHERE子句从数据表中读取数据的通用语法：

```
SELECT column1,cloumn2,...
FROM table_name
WHERE conditon;
```

**参数说明：**

- column1,column2,...是你要选择的列的名称（字段名），如果使用`*`表示选择所有列。

- table_name是你要从中查询数据的表的名称。
- WHERE conditon是用于指定过滤条件的子句。

**更多说明：**

- 查询语句中你可以使用一个或者多个表，表之间使用逗号`，`分隔，并使用WHERE语句设定查询条件。
- 你可以在WHERE子句中指定任何条件。
- 你可以使用AND或者OR指定一个或多个条件。
- WHERE子句也可以运用于SQL的**DELETE**或者**UPDATE**命令。
- WHWER子句类似于程序语言中的if条件，根据MySQL表中的字段来读取指定的数据。

以下未操作符列表，可用于WHERE子句中。

| 操作符 | 描述                                                         | 实例                 |
| :----- | :----------------------------------------------------------- | :------------------- |
| =      | 等号，检测两个值是否相等，如果相等返回true                   | (A = B) 返回false。  |
| <>, != | 不等于，检测两个值是否相等，如果不相等返回true               | (A != B) 返回 true。 |
| >      | 大于号，检测左边的值是否大于右边的值, 如果左边的值大于右边的值返回true | (A > B) 返回false。  |
| <      | 小于号，检测左边的值是否小于右边的值, 如果左边的值小于右边的值返回true | (A < B) 返回 true。  |
| >=     | 大于等于号，检测左边的值是否大于或等于右边的值, 如果左边的值大于或等于右边的值返回true | (A >= B) 返回false。 |
| <=     | 小于等于号，检测左边的值是否小于或等于右边的值, 如果左边的值小于或等于右边的值返回true | (A <= B) 返回 tru    |

**简单实例**

1. 等于条件：

```
SELECT * FROM users WHERE username = 'test';
```

   2.不等于条件：

```
SELECT * FROM users WHERE username != 'runoob';
```

3. 大于条件:

```
SELECT * FROM products WHERE price > 50.00;
```

4. 小于条件:

```
SELECT * FROM orders WHERE order_date < '2023-01-01';
```

5. 大于等于条件:

```
SELECT * FROM employees WHERE salary >= 50000;
```

6. 小于等于条件:

```
SELECT * FROM students WHERE age <= 21;
```

​	7.组合条件（AND、OR）:

```
SELECT * FROM products WHERE category = 'Electronics' AND price > 100.00;
SELECT * FROM orders WHERE order_date >= '2023-01-01' OR total_amount > 1000.00;
```

​	8.模糊匹配条件（LIKE）:

```
SELECT * FROM customers WHERE first_name LIKE 'J%'	//表示匹配以J开头的内容

SELECT * FROM customers WHERE first_name LIKE '%J%'	//表示匹配含J的内容

SELECT * FROM customers WHERE first_name LIKE 'J_'	//表示匹配以J开头且只有两位的内容，匹配占位符 ‘_’

SELECT * FROM customers WHERE first_name LIKE 'J__'	//表示匹配以J开头且只有三位的内容
```

​	9.IN条件：

```
SELECT * FROM countries WHERE country_code IN ('US','CA','MX');
```

​	10.NOT 条件：

```
SELECT * FROM products WHERE NOT category = 'Clothing';
```

​	11.BETWEEN 条件:

```
SELECT * FROM orders WHERE order_date BETWEEN '2023-01-01' AND '2023-12-31';
```

​	12.IS NULL 条件：

```
SELECT * FROM employees WHERE department IS NULL;
```

​	13.IS NOT NULL条件：

```
SELECT * FROM customers WHERE email IS NOT NULL;
```

如果我们想在MySQL数据表中读取指定的数据，WHERE子句是非常有用的。

使用主键作为WHERE子句的条件查询是非常快速的。

如果给定的条件在表中没有任何匹配的记录，那么查询不会返回任何数据。

#### 十一、MySQL更新数据

如果我们需要修改或更新MySQL中的数据，我们可以使用**UPDATE**来操作。

**语法**

以下是UPDATE命令修改MySQL数据表数据的通用SQL语法：

```
UPDARE table_name
SET column1 = value1,column2 = value2,...
WHERE condition;
```

**参数说明：**

- table_name是你要更新数据的表的名称。
- column1,column2,...是你要更新的列的名称。
- value1,value2,...是新的值，用于替换旧的值。
- WHERE condition是一个可选的子句，用于指定更新的行。如果省略WHERE子句，将更新表中的所有行。

**更多说明：**

- 你可以同时更新一个或多个字段。
- 你可以在WHERE子句中指定任何条件。
- 你可以在一个单独表中同时更新数据。

当你需要更细数据表中指定行的数据时WHERE子句是非常有用的。

**实例**

以下实例演示了如何使用UPDATE语句。

1.更新单个列的值：

```
UPDATE employees
SET salary = 60000
WHERE employee_id = 101;
```

2.更新多个列的值：

```
UPDATE orders
SET status = 'Shipped',ship_date = '2023-03-01'
WHERE order_id = 1001;
```

3.使用表达式更新值：

```
UPDATE products
SET price = price * 1.1
WHERE category = 'Electronics';
```

以上SQL语句将每个属于'Electronics'类别的产品的价格都增加了10%。

4.更新符合条件的所有行：

```
UPDATE students
SET status = 'Graduated';
```

以上SQL语句将所有学生的状态更新为'Graduated'。

5.更新使用子查询的值：

```
UPDATE customers
SET total_purchases = (
	SELECT SUM(amount)
	FROM orders
	WHERE orders.costomer_id = customers.customer_id
)
WHERE costomer_type = 'Premium';
```

以上SQL语句通过子查询计算每个'Premium'类型客户的总购买金额，并将该值更新到total_purchases列中。

>注意：在使用UPDATE语句时，请确保你提供足够的田间来确保只有你想要更新的行被修改。如果不提供WHWER子句，将更新表中的所有行，可能导致不可预测的结果。

#### 十二、MySQL删除数据

使用**DELETE FORM**命令来删除MySQL数据表中的记录。

**语法**

以下是**DELETE**语句从MySQL数据表中删除数据的通用语法：

```
DELETE FROM table_name
WHERE condition;
```

**参数说明：**

- table_name是你要删除数据的表的名称。
- WHERE condition是一个可选的子句，用于指定删除的行。**如果省略WHERE子句，将删除表中的所有行。**

**更多说明：**

- 如果没有指定WHERE子句，MySQL表中的所有记录将被删除。
- 你可以在WHERE子句中指定任何条件。
- 你可以在单个表中一次型删除记录。

**实例**

以下实例演示了如何使用DELETE语句。

1.删除符合条件的行：

```
DELETE FROM students
WHERE graduation_year = '2021';
```

2.删除所有行：

```
DELETE FROM students;
```

以上SQL语句删除了orders表中的所有记录，但表结构保持不变。

3.使用子查询删除符合条件的行：

```
DELETE FROM customers
WHERE customer_id IN (
	SELECT customer_id
	FROM orders
	WHERE order_date < '2023-01-01'
);
```

以上SQL语句通过子查询删除了orders表中在'2023-01-01'之前下的订单对应的客户。

> 注意：在使用DELETE语句时，请确保你提供了足够的条件来确保只有你想要删除的行被删除 。如果不提供WHERE子句，将删除表中的所有行，可能导致不可预测的结果。

#### 十三、LIKE 子句

在MySQL中使用**SELECT**命令来读取数据，同时我们可以在**SELECT**语句中使用**WHERE**子句来获取指定的记录。

WHERE子句中是可以使用等号 **=**来设定获取数据的条件，如"runoob_author = 'RUNOOB.COM'"。

但是有时候我们需要获取runoob_author字段含有"COM"字符的所有记录，这时候我们旧需要在WHERE子句中使用**LIKE**子句。

**`LIKE`**子句是MySQL中用于在WHERE子句中进行模糊匹配的关键字。它通常与通配符一起使用，用于搜索符合某种模式的字符串。

 **`LIKE`**子句中使用百分号`%`字符来表示任意字符，类似于UNIX或表达式中的星号`*`。

如果没有使用百分号`%`，**LIKE**子句于等于 `=`效果是一样的。

**语法**

以下是SQL SELECT语句使用LIKE子句从数据表中读取数据的通用语法：

```
SELECT column1,column2,...
FROM table_name
WHERE column_name LIKE pattern;
```

**参数说明：**

- column1,column2,...是你要选择的列的名称，如果使用*表示选择所有列。
- table_name是你要从中查询数据的表的名称。
- column_name是你要应用LIKE子句的列的名称。
- pattern是用于匹配的模式，可以包含通配符。

**更多说明：**

- 你可以在WHERE子句中指定任何条件。
- 你可以在WHERE子句中使用LIKE子句。
- 你可以使用LIKE子句代替等号 `=`。
- LIKE通常与`%`一同使用，类似于一个元字符的搜索。
- 你可以使用AND或者OR指定一个或多个条件。
- 你可以在DELETE或UPDATE命令中使用WHERE...LIKE子句指定条件。

**实例**

以下是一些**`LIKE`**子句的使用实例。

1.百分号通配符%：

`%`通配符表示零个或多个字符。例如，**`'a%'`**匹配以字母**'a'**开头的任何字符串。

```
SELECT * FROM customers WHERE last_name LIKE 'S%';
```

以上SQL语句将选择所有姓氏以'S'开头的客户。

2.下划线通配符`_`：

```
SELECT * FROM products WHERE product_name LIKE '_a%';
```

以上SQL语句将选择产品名称的第二个字符为'a'的所有产品。

3.组合使用%和_:

```
SELECT * FROM user WHERE username LIKE 'a%o_';
```

以上SQL语句将匹配以字母'a'开头，任何是零个或多个字符，接着'o',最后一个是任意字符的字符串，如'aaron'

4.不区分大小写的匹配：

```
SELECT * FROM employees WHERE last_name 'smi%' COLLATE utf8mb4_general_ci;
```

以上SQL语句将选择姓氏以**'smi'**开头的所有员工，不区分大小写。

**LIKE**子句提供了强大的模糊搜索能力，可以根据不同的模式和需求进行定制。在使用时，请确保理解通配符的含义，并根据实际情况进行匹配。

#### 十四、UNION 操作符

**描述**

**MySQL UNION**操作符用于连接两个以上的SELECT语句的结果组合到一个结果集合，**并去除重复的行**。

UNION操作符必须由两个或多个SELECT语句组成，每个SELECT语句的**列数和对应位置的数据类型必须相同**。

**语法**

```
SELECT column1,column2,...
FROM table1
WHERE condition1
UNION
SELECT column1,column2,...
FROM table2
WHERE condition2
[ORDER BY column1,column2,...];
```

**参数说明：**

1.基本的UNION操作：

```
SELECT city FROM customers
UNION
SELECT city FROM suppliers
OEDER BY city;
```

以上SQL语句将选择客户表和供应商表中所有城市的唯一值，并按城市名称升序排序。

2.使用过滤条件的UNION：

```
SELECT product_name FROM products WHERE category = 'Electronics'
UNION
SELECT product_name FROM products WHERE vategory = 'Clothing'
ORDER BY product_name;
```

以上SQL语句将选择电子产品和服装类别的产品名称，并按产品名称升序排序。

3.UNION操作中的列数和数据类型必须相同：

```
SELECT first_name,last_name FROM employees 
UNION
SELECT department_name,NULL FROM departments
ORDER BY first_name;
```

以上SQL语句中，department_name列被映射到了employees表中的first_name列，但是列数和数据类型必须相同。

4.使用UNION ALL不去除重复行：

```
SELECT city FROM customers
UNION ALL
SELECT city FROM suppliers
ORDER BY city;
```

以上SQL语句使用UNION ALL将客户表和供应商表中的所有城市合并在一起，不去除重复行。

> UNION操作符在合并结果集时会去除重复行，而UNION ALL不会去除重复行，因此UNION ALL的性能可能更好，但如果你确实希望去除重复行，可以使用UNION。

#### 十五、ORDER BY(排序)语句

我们知道从MySQL表中使用**SELECT**语句来读取数据。

如果我们需要对读取的数据进行排序，我们就可以使用MySQL的**ORDER BY**子句来设定你想按哪个字段哪种方式进行排序，再返回搜索结果。

**ORDER BY(排序)**语句可以按照一个或多个列的值进行升序（**ASC**）或降序（**DESC**）排序。

**语法**

```
SELECT column1,column2,...
FROM table_name
ORDER BY column1 [ASC|DESC],column2 [ASC,DESC],...;
```

**参数说明：**

- column1，column2，...是你要选择的列的名称，如果使用*表示选择所有列。
- table_name是你要从中查询数据的表的名称。
- ORDER BY column1 [ASC | DESC],column2 [ASC |DESC],...是用于指定排序顺序的子句。ASC表示升序（默认），DESC表示降序。

**更多说明：**

- 你可以使用任何字段作为排序的条件，从而返回排序后的查询结构。
- 你可以设定多个字段来排序。
- 你可以使用ASC或DESC关键字来设置查询结构式升序或降序排列。默认情况下，它是按升序排列。

**实例**

1.单列排序：

```
SELECT * FROM pruducts
ORDER BY prodect_name ASC;
```

以上SQL语句将选择产品表products中的所有产品，并按产品名称升序ASC排序。

2.多列排序：

```
SELECT * FROM employees
ORDER BY department_id ASC,hire_date DESC;
```

以上SQL语句将选择员工表employees中的所有员工，并先按部门ID升序ASC排序，任何再相同部门中按雇佣日期降序DESC排序。

3.使用数字表示列的位置：

```
SELECT first_name,last_name,salary
FROM employees
OEDER BY 3 DESC,1 ASC;
```

以上SQL语句将选择员工表employees中的名字和工资列，并按第三列（salary）降序DESC排序，任何按第一列（first_name）升序ASC排序。

4.使用表达式排序：

```
SELECT product_name,price * discount_rate AS discounted_price
FROM products
ORDER BY discounted_price DESC;
```

以上SQL语句将选择产品表products中的产品名称和根据折扣率计算折扣后的价格，并按折扣后价格降序DESC排序。

5.使用NULLS FIRST或NULLS LAST处理NULL值：

```
SELECT product_name,price
FROM products
ORDER BY price DESC NULLS LAST;
```

以上SQL语句将选择产品表products中的产品名称和价格，并按价格降序排序后，将NULL值排在最后。

> ORDER BY子句是一个强大的工具，可以根据不同的业务需求查询结果进行排序。在实际应用中，注意选择适当的列和排序顺序，以获得符合期望的排序效果。

MySQL 排序我们知道从 MySQL 表中使用 SQL SELECT 语句来读取：

**MySQL 拼音排序**

如果字符集采用的是 gbk(汉字编码字符集)，直接在查询语句后边添加 ORDER BY：

```
SELECT * 
FROM runoob_tbl
ORDER BY runoob_title;
```

如果字符集采用的是 utf8(万国码)，需要先对字段进行转码然后排序：

```
SELECT * 
FROM runoob_tbl
ORDER BY CONVERT(runoob_title using gbk);
```

#### 十六、GROUP BY语句

GROUP BY语句根据一个或多个列对结果集进行分组。在分组的列上我们可以使用COUNT，SUM，AVG等函数。

GROUP BY语句是SQL查询中用于汇总和分析数据的重要工具，尤其在处理大量数据时，它能够提供有用的汇总信息。

**GROUP BY语法**

```
SELECT column1，aggregate_function(column2)
FROM table_name
WHERE codition
GROUP BY column1;
```

- column1:指定分组的列。
- aggregate_function(column2):对分组后的每个执行的聚合函数。
- table_name:要查询的表名。
- condition:可选，用于筛选结果的条件。

假设有一名orders的表，包含以下列：order_id、customer_id、order_rate和order_amount。

我们想要按照customer_id进行分组，并计算每个客户的订单总金额，SQL语句如下：

```
SELECT customer_id,SUM(order_amount) AS total_amount
FROM orders
GROUP BY customer_id;
```

以上实例中，我们使用GROUP BY customer_id将结果按customer_id列分组，然后使用SUM(order_amount)计算每个组中order_amount列的总和。

AS total_amount是为了给计算结果取一个别名，使查询结构更易读。

**注意事项：**

- GROUP BY子句通常与聚合函数一起使用，因为分组后需要对每个组进行聚合操作。
- SELECT子句中的列通常要么使分组列，要么使聚合函数的参数。
- 可以使用多个列进行分组，只需在GROUP BY子句中使用逗号分隔列名即可。

```
SELECT column1, column2, aggregate_function(column3)
FROM TABLE_NAME
WHERE condition
GROUP BY column1, column2;
```

#### 十七、JOIN语句

在前几章节中，我们以及学会了如何在一张表中读取数据，这是相对简单的，但是真正的应用中经常需要从多个数据表中读取数据。

本章节将介绍如何使用MySQL的**JOIN**在两个或多个表中查询数据。

可以在SELECT、UPDATE和DELETE语句中使用MySQL的JOIN来联合多表查询。

**JOIN按照功能大致分为如下三类：**

- **INNER JOIN(内连接，或等值连接)：**获取两个表中字段匹配关系的记录。
- **LEFT JOIN(左连接)：**获取左表所有记录，即使右表没有对应匹配的记录。
- **RIGHT JOIN(右连接)：**与LEFE JOIN相反，用于获取右表所有记录，即使左表没有对应匹配的记录。

**INNER JOIN**

INNER JOIN返回两个表中满足连接条件的匹配行，以下是INNER JOIN语句的基本语法：

```
SELECT column1，column2，...
FROM tabel1
INNER JOIN table2 ON table1.column_name = table.column_name;
```

**参数说明：**

- column1,column2,...是你要选择的列的名称，如果使用*表示选择所有列。
- table1,table2是要连接的两个表的名称。
- table.column_name = table2.column_name是连接条件，指定了两个表中用于匹配的列。

1.简单的INNER JOIN:

```
SELECT orders.order_id,customers.customer_name
FROM oders
INNER JOIN customers ON orders.customer_id = customers.customer_id;
```

以上SQL语句将选择orders表和customes表中满足连接条件的订单ID和客户名称。

2.使用表别名：

```
SELECT o.order_id,c.customer_name
FROM orders AS o
INNER JOIN customers AS c ON o.customer_id = c.customer_id;
```

以上SQL语句使用表别名o和c作为orders和customers表的别名。

3.多表INNER JOIN:

```
SELECT orders.order_id,customers.customer_name,products.product_name
FROM orders
INNER JOIN customers ON orders.customer_id = customers.customer_id
INNER JOIN order_items ON orders.order_id = order_items.order_id
INNER JOIN prodcuts ON order_items.product_id = products.product_id;
```

以上SQL语句涉及了orders、customers、order_itmes和products四个表的连接。它选择了订单ID、客户名称和产品名称，连接了这些表的关联列。

4.使用WHERE子句进行过滤：

```
SELECT orders.order_id,customers.customer_name
FROM orders
INNER JOIN customers ON orders.customer_id = customers.customer_id
WHERE orders.order_date > '2023-01-01';
```

以上SQL语句在INNER JOIN后使用WHERE子句，过滤了订单日期在'2023-01-01'及以后的订单。

**LEFT JOIN**

LEFT JOIN返回左表的所有行，并包括右表中没有匹配的行，如果右表中没有匹配的行，将返回NULL值，以下是LEFR JOIN语句的基本语法：

```
SELECT column1,column2,...
FROM table1
LEFT JOIN table2 ON tabel1.colmn_name = table2.cloumn_name;
```

1.简单的LEFT JOIN：

```
SELECT customers.customer_id,coustomers.customer_name,orders.order_id
FROM customers
LEFT JOIN orders ON customers.customer_id = orders.customer_id;
```

以上SQL语句将选择客户表中的客户ID和客户名称，并包括左表customers中的所有行，以及匹配的订单ID（如果有的话）。

2.使用表别名：

```
SELECT c.customer_id,c.customer_name,o.order_id
FROM customers AS c
LEFT JOIN orders AS o ON c.customer_id = o.customer_id;
```

以上SQL语句使用表别名c和o分别代替customers和orders表的名称。

3.多表LEFT JOIN：

```
SELECT customers.customer_id,customers.customer_name,orders.order_id,products.product_name
FROM customers
LEFT JOIN orders ON customers.customer_id = orders.customer_id
LEFT JOIN order_items ON orders.order_id = order_items.order_id
LEFT JOIN products ON order_itmes.product_id = products.product_id;
```

以上SQL语句连接了customers、orders、order_items和products四个表，并选择了客户ID、客户名称、订单ID和产品名称。

左连接保证了即使在order_item或products中没有匹配的行，仍然会返回客户和订单的信息。

4.使用WHERE子句进行过滤：

```
SELECT customers.customer_id,customers.customer_name,orders.order_id
FROM customers
LEFT JOIN orders ON customers.customer_id = orders.customer_id
WHERE orders.order_date >= '2023-01-01' OR orders.order_id IS NULL;
```

以上SQL语句在LEFT JOIN后使用WHERE子句，过滤了订单日期在'2023-01-01'及以后的订单，以及没有匹配订单的客户。

LEFT JOIN是一张常用的连接类型，尤其在需要返回左表所有行的情况下。但右表中没有匹配的行时，相关列将显示为NULL。

在使用LEFT JOIN时，请确保理解连接条件并根据需求过滤结果。

**RIGHT JOIN**

RIGHR JOIN返回右表的所有行，并包括左表中匹配的行，如果左表中没有匹配的行，将返回NULL值，以下时RIGHT JOIN语句的基本语法：

```
SELECT column1,column2,...
FROM table1
RIGHT JOIN table2 ON table1.column_name = table2.column_name;
```

以下是一个简单的RIGHT JOIN实例：

```
SELECT customers.customer_id,orders.order_id
FROM customers
RIGHT JOIN orders ON customers.customer_id = orders.customer_id;
```

以上SQL语句将选择右表orders中的所有订单ID，并包括左表customers中匹配的客户ID。如果在customers中没有匹配的客户ID，相关列将显示为NULL。

在开发过程中，RIGHT JOIN并不经常使用，因为它可以用LEFT JOIN和表的顺序交换来实现相同的效果。例如，上， 的查询可以通过使用LEFT JOIN改写为：

```
SELECT customers.customer_id,oders.order_id
FROM orders
LEFT JOIN customers ON orders.customer_id = customers.customer_id;
```

以上SQL语句返回相同的结果，因为LEFT JOIN于RIGHE JOIN是对称的。在实际使用中，你可以根据个人偏好或组织规范选择使用哪种形式。

#### 十八、NULL值处理

我们已经知道MySQL使用**SELECT**命令及**WHERE**子句来读取数据表中的数据，但是当提供的查询条件字段为NULL时，改命令keneng无法正常工作。

在MySQL中，NULL用于表示缺失的或未知的数据，处理NULL值需要特别小心，因为在数据库中它可能会导致不同于预期的结果。

为了处理这种情况，MySQL提供了三大运算符：

- **IS NULL**：当列的值是NULL，此运算符返回true。
- **IS NOT NULL**：当列的值不为NULL，运算符返回true。
- **<=>**：比较操作符（不同于 = 运算符），当比较的两个值相等或者都为NULL时返回true。

关于NULL的条件比较运算时比较特殊的，你不能使用 = NULL或 != NULL在列中查找NULL值。

在MySQL中，**NULL值于任何其他值比较（即使是NULL）永远返回NULL，即NULL = NULL返回NULL**。

MySQL中处理NULL使用IS NULL 和IS NOT NULL运算符。

> 注意：
>
> SELECT  *,columnName1 + ifnull(columnName2,0) FROM tablename;
>
> 
>
> columnName1,columnName2为int型，当columnName2中，有值为NULL时，columnName1 + columnName2 = NUll,ifnull(columnName2,0)把columnName2中的null值转为0。

**MySQL中处理NULL值的常见注意事项和技巧**

1.检查是否为NULL：

要检查某列是否为NULL，可以使用IS NULL 或 IS NOT NULL条件。

```
SELECT * FROM employees WHERE department_id IS NULL;
SELECT * FROM employees WHERE department_id IS NOT NULL;
```

2.使用COALESCE函数处理NULL：

```
SELECT product_name,COALESCE(stock_quantity,0) AS actual_quantity
FROM products;
```

以上SQL语句中，如果stock_quantity列为NULL，则COALESCE将返回0。

3.使用IFNULL函数处理NULL：

IFNULL函数是COALESCE的MySQL特定版本，它接受两个参数，如果第一个参数为NULL，则返回第二个参数。

```
SELECT product_name,IFNULL(stock_quantity,0) AS actual_quantity
FROM products;
```

4.NULL排序：

在使用ORDER BY子句进行排序时，NULL值默认会被放在排序的后面。如果希望将NULL值放在最前面，可以使用ORDER BY column_name ASC NULLS FIRST,反之使用ORDER BY column_name DESC NULLS LAST。

```
SELECT product_name,price
FROM products
ORDER BY price ASC NULLS FIRST;
```

5.使用<=>操作符进行NULL比较：

<=>操作符是MySQL中用于比较两个表达是否相等的特殊操作符，对于NULL值的比较也会返回TRUE。它可以用于处理NULL值的等值比较。

```
SELECT * FROM employees WHERE commission <=> NULL;
```

6.注意聚合函数对NULL的处理：

在使用聚合函数（如COUNT，SUM，AVG）时，它们会忽略NULL值，因此可能会得到不同预期的结果。如果希望将NULL视为0，可以使用COALESCE或IFNULL。

```
SELECT AVG(COALESCE(salary,0)) AS avg_salary FROM emplpyees;
```

这样即使salary为NULL，聚合函数也会将其视为0.

> 处理NULL时，要特别小心确保查询和操作的语义符合预期。在设计表结构时，也需要考虑NULL值的使用场景和合理性。

#### 十九、MySQL正则表达式

在前面章节我们已经了解到MySQL可以通过LIKE...%来进行模糊匹配。

MySQL同样也支持其他正则表达式的匹配，MySQL中使用**REGEXP**和**RLIKE**操作符来进行正则表达式匹配。

下表中的正则模式可应用于REGEXP操作符中。

| 模式       | 描述                                                         |
| :--------- | :----------------------------------------------------------- |
| ^          | 匹配输入字符串的开始位置。如果设置了 RegExp 对象的 Multiline 属性，^ 也匹配 '\n' 或 '\r' 之后的位置。 |
| $          | 匹配输入字符串的结束位置。如果设置了RegExp 对象的 Multiline 属性，$ 也匹配 '\n' 或 '\r' 之前的位置。 |
| .          | 匹配除 "\n" 之外的任何单个字符。要匹配包括 '\n' 在内的任何字符，请使用像 '[.\n]' 的模式。 |
| [...]      | 字符集合。匹配所包含的任意一个字符。例如， '[abc]' 可以匹配 "plain" 中的 'a'。 |
| [^...]     | 负值字符集合。匹配未包含的任意字符。例如， '[^abc]' 可以匹配 "plain" 中的'p'。 |
| p1\|p2\|p3 | 匹配 p1 或 p2 或 p3。例如，'z\|food' 能匹配 "z" 或 "food"。'(z\|f)ood' 则匹配 "zood" 或 "food"。 |
| *          | 匹配前面的子表达式零次或多次。例如，zo* 能匹配 "z" 以及 "zoo"。* 等价于{0,}。 |
| +          | 匹配前面的子表达式一次或多次。例如，'zo+' 能匹配 "zo" 以及 "zoo"，但不能匹配 "z"。+ 等价于 {1,}。 |
| {n}        | n 是一个非负整数。匹配确定的 n 次。例如，'o{2}' 不能匹配 "Bob" 中的 'o'，但是能匹配 "food" 中的两个 o。 |
| {n,m}      | m 和 n 均为非负整数，其中n <= m。最少匹配 n 次且最多匹配 m 次。 |

**正则表达式匹配的字符类**

- `.`：匹配任意单个字符。
- `^`：匹配字符串的开始。
- `$`：匹配字符串的结束。
- `*`：匹配零个或多个前面的元素。
- `+`：匹配一个或多个前面的元素。
- `?`：匹配零个或一个前面的元素。
- `[abc]`：匹配字符集中的任意一个字符。
- `[^abc]`：匹配除了字符集中的任意一个字符以外的字符。
- `[a-z]`：匹配范围内的任意一个小写字母。
- `\d`：匹配一个数字字符。
- `\w`：匹配一个字母数字字符（包括下划线）。
- `\s`：匹配一个空白字符。

****

**使用REGEXP进行模式匹配**

REGEXP是用于进行正则表达式匹配的运算符。

REGEXP用于检查一个字符串是否匹配指定的正则表达式模式，以下是REGEXP运算符的基本语法：

```
SELECT column1,column2,...
FROM table_name
WHERE columne_name REGEXP 'pattern';
```

**参数说明：**

- column1,column2,...是你要选择的列的名称，如果使用*表示选择所有列。
- table_name是你要从中查询数据的表的名称。
- column_name是你要进行正则表达式匹配的列的名称。
- 'pattern'是一个正则表达式模式。

1.查找name字段中以'st'为开头的所有数据：

```
SELECT name FROM person_tab1 WHERE name REGEXP '^st';
```

2.查找name字段中以'ok'为结尾的所有数据：

```
SELECT name FROM person_tb1 WHERE name REGEXP 'ok$';
```

3.查找name字段中包含'mar'字符串的所有数据：

```
SELECT name FROM person_tb1 WHERE name REGEXP 'mar';
```

4.选择订单表中描述包含'item'后跟一个或多个数字的记录。

```
SELECT * FROM orders WHERE order_description REGEXP 'item[0-9]+';
```

5.使用**BINARY**关键字，使得匹配区分大小写：

```
SELECT * FROM products WHERE product_name REGEXP BINARY 'apple';
```

6.使用OR进行多个匹配条件，以下将选择形式为'Smith'或'Johnson'的员工记录：

```
SELECT * FROM employees WHERE last_name REGEXP 'Smith|Johnson';
```

**使用RLIKE进行模式匹配**

RLIKE是MySQL中用于进行正则表达式匹配的运算符，与REGEXP是一样的，RLIKE和REGEXP可以互换使用，没有区别。

以下是使用RLIKE进行正则表达式匹配的基本语法：

```
SELECT column1,column2,...
FROM table_name
WHERE column_name RLIKE 'pattern';
```

**参数说明：**

- column1，column2，...是你要选择的列的名称，如果使用*表示所有列。
- table_name是你要从中查询数据的表的名称。
- column_name是你要进行正则表达式匹配的列的名称。
- 'pattern'是一个正则表达式模式。

```
SELECT * FROM products WHERE product_name RLIKE '^[0-9]';
```

以上SQL语句选择产品名称以数字开头的所有产品。

#### 二十、MySQL事务

MySQL事务主要用于处理操作量大，复杂度高的数据。比如说，在人员管理系统中，你删除一个人员，你既需要删除人员的基本资料，也要删除和该人员相关的信息，如信箱，文章等等，这样，这些数据库操作语句就构成一个事务！

在MySQL中，事务是一组SQL语句的执行，它们被视为一个单独的工作单元。

- 在MySQL中只用使用了Innodb数据库引擎的数据库或表才支持事务。
- 事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行。
- 事务用来管理**insert、update、delete**语句

一般来说，事务是必须满足4个条件（ACID）：：原子性（Atomicity，或不可分隔性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）。

- **原子性**：一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。
- **一致性**：在事务开始先和事务结束以后，数据库的完成性没有被破环。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精准度、串联型以及后数据库可以自发性地完成预定的工作。

- **隔离性**：数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行到时数据的不一致。事务格力分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable）和串行化（Serializable）。
- **持久性**：事务处理结束后，对数据的修改就是永久的，即使系统故障也不会丢失。

> 在MySQL命令行的默认设置下，事务都是自动提交的，即执行SQL语句后就会马上COMMIT操作。因此要显式地开启一个事务必须使用命令BEGIN或STAR TRANSACTION，或者执行命令SET AUTOCOMMIT = 0，用来禁止使用当前会话的自动提交。

**事务控制语句：**

- BEGIN或START TRANSACTION显式地开启一个事务；
- COMMIT也可以使用COMMIT WORK，不过二者是等价的。COMMIT回提交事务，并使已对数据库进行的所有修改成为永久性的；
- ROLLBACK也可以使用ROLLBACK WORK，不过二者是等价的。回滚回结束用户的事务，并撤销正在进行的所有未提交的修改；
- SACEPOINT identifier,SAVEPOINT允许再事务中创建一个保存点，一个事务中可以有多个SACEPOINT ；
- RELEASE SAVEPOINT identifier删除一个事务的保存点，当没有指定的保存点时，执行该语句回抛出一个异常；
- ROLLBACK TO identifier把事务回滚到标记点；
- SET TANSACTION用来设置事务的隔离级别。InnoDB存储引擎提供事务的隔离级别。InnoDB存储引擎提供事务的隔离级别有读未提交(READ UNCOMMITTED)、READ COMMITTED、REPEATABLE READ和SERILIZABLE。

![image-20240409112406872](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240409112406872.png)

**MySQL事务处理主要有两种方法：**

1.用BEGIN，ROLLBSCK，COMMIT来实现

- **BEGIN或START TRANSACTION**：用于开始一个事务。
- **ROLLBACK**事务回滚，取消之前的更改。
- **COMMIT：**事务确认，提交事务，使更改永久生效。

2.直接用SET来改变MySQL的自动提交模式：

- **SET AUTOCOMMIT = 0**禁止自动提交
- **SET AUTOCOMMIT = 1**开启自动提交

**BEGIN或START TRANSACTION**-----用于开始一个事务：

```
BEGIN; -- 或者使用START TRANSACTION;
```

**COMMIT** -- 用于提交事务，将所有的修改永久保存到数据库：

```
COMMIT;
```

**ROLLBACK** -- 用于回滚事务，撤销自上次提交以来所做的所有更改：

```
ROLLBACK;
```

**SAVEPOINT**-- 用于再事务中设置保存点，以便稍后能回滚到该点：

```
SAVEPOINT savepoint_name;
```

**ROLLBACK TO SAVEPOINT** -- 用于回滚到之前设置的保存点：

```
ROLLBACK TO SAVEPOINT savepoint_name；
```

下面使一个简单的MySQL事务的例子：

```
-- 开启事务
START TRANSACTION;

-- 执行一些SQL语句
UPDATE accounts SET balance = balance - 100 WHERE user_id = 1;
UPDATE accounts SET balance = balance + 100 WHERE user_id = 2;

-- 判断是否提交还是回滚
IF (条件) THEN
	COMMIT; -- 提交事务
ELSE
	ROLLBACK;
END IF;
```

#### 二十一、ALTER命令

我们需要修改数据表名或者修改数据表字段时，就需要使用到**ALTER**命令。

MySQL的**ALTER**命令用于修改数据库、表和索引等对象的结构。

**ALTER**命令允许你添加、修改或删除数据库对象，并且可以 用于更改表的列定义、添加约束、创建和删除索引等操作。

**ALTER**命令非常强大，可以再数据库结构发生变化时进行灵活的修改和调整。

以下是**ALTER**命令的常见用法和实例：

**1.添加列**

```
ALTER TABLE table_name
ADD COLUMN new_column_name datatype;
```

以下SQL语句在employees表中添加了一个名为birth_date的日期列：

```
ALTER TABLE employees
ADD COLUMN birth_date DATE;
```

**2.修改列的数据类型**

```
ALTER TABLE table_name
MODIFY COLUMN column_name new_datatype;
```

以下SQL语句将employees表中的salary列的数据类型修改为DECIMAL(10,2):

```
ALTER TABLE employees
MODIFY COLUMN salary DECIMAL(10,2);
```

**3.修改列名**

```
ALTER TABLE table_name
CHANGE COLUMN old_column_name new_column_name datatype;
```

以下SQL语句将employees表中的某个列名字有name改为user_name,并且可以同时修改数据类型：

```
ALTER TABLE employees 
CHANGE COLUMN name user_name VARCHAR(255);
```

**4.删除列**

```
ALTER TABLE table_name
DROP COLUMN column_name;
```

以下SQL语句将employees表中的birth_date列删除：

```
ALTER TABLE employees
DORP COLUMN birth_date;
```

**5.添加PRIMARY KEY**

```
ALTER table_name
ADD PRIMARY KEY (column_name);
```

**6.添加ROREIGN KEY**

```
ALTER TABLE child_table
ADD CONSTRAINT fk_name
FOREIGN KEY (column_name)
REFERENCES parent_table (column_name);
```

以下SQL语句在orders表中添加了一个外键，关联到customers表的customer_id列：

```
ALTER TABLE orders
ADD CONSTRAINT fk_customer
FOREING KEY (customer_id)
REFERENCES customers (customer_id);
```

**7.修改表名**

```
ALTER TABLE old_table_name
RENAME TO new_table_name;
```

以下SQL语句将表名由employees修改为staff：

```
ALTER employees
RENAME TO staff;
```

> 注意：
>
> 但在使用ALTER命令时要格外小心，因为一些操作可能需要重建表或索引，这可能回影响数据库的性能和允许时间。
>
> 在进行重要的结构修改时，建议先备份数据，并在生产环境中谨慎操作。

#### 二十二、MySQL索引

MySQL索引是一种数据结构，用于加快数据库查询的速度和性能。

MySQL索引的建立对于MySQL的高效运行是很重要的，索引可以大大提高MySQL的检索速度。

> MySQL索引类似于书籍的索引，通过存储指向数据行的指针，可以快速定位和方位表中的特定数据。
>
> 打个比方，如果合理的设计且使用索引的MySQL是一辆兰博基尼的话，那么没有设计使用索引的MySQL就是一个人力三轮车。
>
> 拿汉语字典的目录页（索引）打比方，我们可以按拼音、笔画、偏旁部首等排序的目录（索引）快速查找到需要的字。

索引分单列索引和组合索引：

- **单列索引**：即一个索引只包含单个列，一个表可以由多个单列索引。
- **组合索引**：即一个索引包含多个列。

创建索引时，你需要确保该索引是应用在SQL查询语句的条件（一般作为WHERE子句的条件）。

实际上，索引也是一张表表，**该表保存了主键和索引字段，并指向实体表的记录。**

索引虽然能够提高查询性能，但也需要注意以下几点：

- 索引需要占用额外的存储空间
- 对表大插入、更新和删除操作时，索引需要维护，可能会影响性能。
- 过多或不合理的索引可能会导致性能下降，因此需要谨慎选择和规划索引。

**普通索引**

索引能够显著提高查询的速度，尤其是在大型表中进行搜索时。通过使用索引，MySQL可以直接定位到满足查询条件的数据行，而无需逐行扫描整个表。

**创建索引**

使用**CREATE INDEX**语句可以创建普通索引。

普通索引是最常见的索引类型，用于加速对表中数据的查询。

**CREATE INDEX**的语法：

```
CREATE INDEX index_name
ON table_name (column1 [ASC|DESC],column2 [ASC|DESC]， ...)；
```

**参数说明：**

- CREATE INDEX:用于创建普通索引的关键字。
- index_name：指定要创建索引的名称。索引名称在表中必须是唯一的。
- table_name：指定要在哪个表上创建索引。
- (column1,column2,...)：指定要索引的表列名。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
- ASC和DESC（可选）：用于指定索引的排序顺序。默认情况下，索引以升序（ASC）排序。

以下实例假设我们由一名为students的表，包含id、name和age列，我们将在name列上创建一个普通索引。

```
CREATE INDEX ind_name ON students (name);
```

上述语句将在students表的name列上创建一个名为idx_name的普通索引，这将有助于提高通过姓名进行搜索的查询性能。需要注意的是，如果表中的数据量较大，索引的创建可能会花费一些时间，但一旦创建完成，查询性能将会显著提高。

**修改表结构（添加索引）**

我们可以使用**ALTER TABLE**命令可以在已有的表中创建索引。

ALTER TABLE 允许你修改表的结构，包括添加、修改或删除索引。

ALTER TABLE创建索引的语法：

```
ALTER TABLE table_name
ADD INDEX index_name (column1 [ASC|DESC],column2 [ASC|DESC],...);
```

**参数说明：**

- ALTER TABLE：用于修改表结构的关键字。
- table_name：指定要修改的表的名称。
- ADD INDEX：添加索引的子句。ADD INDEX用于创建普通索引。
- index_name：指定要创建的索引的名称。索引名称在表中必须是唯一的。
- (column1.column2,...)：指定要索引的表列名。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
- ASC和DESC（可选）：用于指定索引的排序顺序，默认情况下，索引以升序（ASC）排序。

下面是一个实例，我们将在已存在的名为employees的表上创建一个普通索引：

```
ALTER TABLE employees
ADD INDEX idx_age (age);
```

上述语句将在employees表中的age列上创建一个名为idx_age的谱索引。

**创建表的时候直接指定**

我们可以在创建表的时候，在**CREATE TABLE**语句中直接指定索引，以创建表和索引的组合。

```
CERATE TABLE table_name (
	column1 data_type,
	column2 data_type,
	...,
	INDEX index_name (column1 [ASC|DESC],cloumn2 [ASC|DESC],...)
);
```

**参数说明：**

- CREATE TABLE：用于创建新表的关键字。
- table_name：指定要创建的表的名称。
- column1,column2,...：定义表的列名和数据类型。你可以指定一个或者多个列作为索引组合。这些列的数据类型通常是数值、文本或日期。
- INDEX：用于创建普通索引的关键字。
- index_name：指定要创建的索引的名称。索引名称在表中必须是唯一的。
- (column1,column2,...)：指定要索引的表列名。可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
- ASC和DESC（可选）：用于指定索引的排序顺序。默认情况下，索引以升序（ASC）排序。

下面是一个实例，我们要创建一个名为students的表，并在age列上创建一个普通索引。

```
CREATE TABLE students (
	id INT PRIMARY KEY,
	name VARCHAR(50),
	age INT,
	INDEX idx_age(age)
);
```

**删除索引的语法**

我们可以使用DROP INDEX语句来删除索引。

**DROP INDEX语法：**

```
DROP INDEX index_name ON table_name;
```

- DROP INDEX：用于删除索引的关键字。
- index_name：指定要删除的索引的名称。
- ON table_name：指定要在哪个表上删除索引。

使用**ALTER TABLE**语句删除索引的语法如下：

```
ALTER TABLE table_name
DROP INDEX index_name;
```

- ALTER TABLE：用于修改表结构的关键字。
- table_name：指定要修改的表的名称。
- DROP INDEX：用于删除索引的子句。
- index_name：指定要删除的索引的名称。

以下实例假设我们有一名名为employees的表，并在age列上有一名为idx_age的索引，选择我们要删除这个索引：

```
DROP INDEX idx_age ON employees;
```

或使用ALTER INDEX语句:

```
ALTER TABLE employees
DROP INDEX idx_age;
```

这两个命令都会从employees表中删除名为idx_age的索引。

如果索引不存在，执行命令时会产生错误，因此，在删除索引之前最好确认该索引是否存在，或者使用错误处理机制来处理可能的错误情况。

****

**唯一索引**

在MySQL中，你可以使用**CREATE UNIQUE INDEX**语句来创建唯一索引。

唯一索引确保索引中的值时唯一的，不允许有重复值。

**创建索引**

```
CREATE UNIQUE INDEX index_name
ON table_name (column1 [ASC|DESC],column2 [ASC|DESC],...);
```

- CREATE UNIQUE INDEX：用于创建唯一索引的关键字组合。
- index_name：指定要创建的唯一索引的名称。索引名称在表中必须时唯一的。
- rable_name：指定要在哪个表上创建唯一索引。
- (column1,column2,...)：指定要索引的表列名。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
- ASC和DESC（可选）：用于指定索引的排序顺序。默认情况下，索引以升序（ASC）排序。

以下是一个创建唯一索引的实例：假设我们有一个名为employees的表，包含id和email列，选择我们想在email列上创建一个唯一索引，以确保每个员工的电子邮箱地址都是唯一的。

```
CREATE UNIQUE INDEX idx_email ON employees (email);
```

以上实例将在employees表的email列上创建一个名为idx_email的唯一索引。

**修改表结构添加索引**

我们可以使用**ALTER TABLE**命令来创建唯一索引。

ALTER TABLE命令允许你修改以及存在的表结构，包括添加新的索引。

```
ALTER TABLE table_name
ADD CONSTRAINT unique_constraint_name UNIQUE (column1,column2,...);
```

- ALTER TABLE：用于修改表结构的关键字。
- table_name：指定要修改的表的名称。
- ADD CONSTAINT：这是用于添加约束（包括唯一索引）的关键字。
- unique_constraint_name：指定要创建的唯一索引的名称，约束名称在表中必须是唯一的。
- UNIQUE (column1,column2,...)：指定要索引的表列名。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。

以下是一个使用ALTER TABLE命令创建唯一索引的实例：假设我们有一个名为employees的表，闹含id和email列，现在我们想在email列上创建一个唯一索引，以确保每个员工的电子邮件地址都是唯一的。

```
ALTER TABLE employees
ADD CONTRAINT idx_email UNIQUE (email);
```

以上实例将在employees表的email列上创建一个名为idx_email的唯一索引。

> 注意：
>
> 如果表中已经有重复的email值，那么添加唯一的索引将会失败。在创建索引之前，你可以需要确保表中的email列没有重复的值。

**创建表的时候直接指定**

我们也可以在创建表的同时，你可以在**CREATE TABLE**语句中使用**UNIQUE**关键字来创建唯一索引。

这将在表创建时同时定义唯一索引约束。

**CREATE TABLE**语句中创建唯一索引的语法：

```
CREATE TABLE table_name (
	column1 data_type,
	column2 date_type,
	...,
	CINSTRAINT index_name UNIQUE (column1 [ASC|DESC],column2 [ASC|DESC],...)
);
```

- CREATE TABLE：用于创建新表的关键字。
- table_name：指定要创建的表的名称。
- (column1,column2,...):定于表的列名和数据类型。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
- CONSTEAINT：用于添加约束的关键字。
- index_name：指定要创建唯一索引的名称。约束名称在表中必须是唯一的。
- UNIQUE (column1,column2,...)：指定要索引的表列名。

以下是一个创建表时创建唯一索引的实例：假设我们要创建一个名为employees的表，其中包含id、name和email列，我们希望email列的值是唯一的，因此我们要在创建表时定义唯一索引。

```
CREATE TABLE employees (
	id INT PRIMARY KEY,
	name VARCHAR(50),
	email VARCHAR(100) UNIQYUE
);
```

在这个例子中，email列被定义为唯一索引，因为在它的后加上了UNIQUE关键字。

使用**UNIQUE**关键字后，索引名称将自动生成，你也可以根据需要指定索引名称。

****

**使用ALTER命令添加和删除索引**

有四种方式来添加数据表索引：

- **ALTER TABLE tbl_name ADD PRIMARY KEY (column_list)**：该语句添加一个主键，这意味着索引值必须是唯一的，且不能为NULL。
- **ALEER TABLE tbl_name ADD UNIQUE index_name (column_list)**：这条语句创建索引的值必须是唯一的(除了NULL外，NULL可能会出现多次)。
- **ALTER TABLE tbl_name ADD INDEX index_name (column_list)**：添加普通索引，索引值可出现多次。
- **ALTER TABLE tbl_name ADD FULLTEXT index_name (column_list)**：该语句指定了索引为FULLTEXT，用于全文索引。

以下实例为表中添加索引。

```
ALTER TABLE testlter_ble ADD index (c);
```

你还可以在ALTER命令中使用DROP子句来删除索引。

**使用ALTER命令添加和删除主键**

主键作用于列上(可以一个列或多个列联合主键)，添加主键索引时，你需要确保该主键默认不为空（NOT NULL）。实例如下：

```
ALTER TABLE table_name MODIFY i INT NOT NULL;
ALTER TABLE banle_name ADD PRIMARY KEY (i);
```

你也可以使用ALTER命令删除主键：

```
ALTER TABLE table_name DROP PRIMARY KEY;
```

**显示索引信息**

可以使用**SHOW INDEX**命令来列出表中的相关的索引信息。

可以通过添加**\G**来格式化输出信息。

**SHOW INDEX**语句：

```
SHOW INDEX FROM table_name\G
```

- SHOW INDEX：用于显示索引信息的关键字。
- FROM table_name：指定要查看索引信息的表的名称。
- \G:格式化输出信息。

执行上述命令后，将会显示指定表中索引的详细信息，包括索引名称（Key_name）、索引列（Column_name）、是否是唯一索引（Non_unique）、排序方式（Collation）、索引的基数（Cardinality）等。

#### 二十三、MySQL临时表

MySQL临时表在我们需要保存一些临时数据时时非常有用的。

临时表只在当前连接可见，当关闭连接时，MySQL会自动删除表并释放所有空间。

在MySQL中，临时表时一种在当前绘画中存在的表，它在绘画结束时会自动被销毁。

如果你使用了其他MySQL客户端程序连接MySQL数据库服务器来创建临时表，那么只有当关闭客户端程序时才会销毁临时表，当然也可以手动销毁。

**创建临时表**

```
CREATE TEMPORARY TABLE temp_table_name (
  column1 datatype,
  column2 datatype,
  ...
);
```

或简写为：

```
CREATE TEMPORARY TABLE tempble_name AS SELECT column1,column2,...
FROM source_table
WHERE condition;
```

**插入数据到临时表**

```
INSERT INTO temp_table_name (column1, column2, ...)
VALUES (value1, value2, ...);
```

**查询临时表**

```
SELECT * FROM temp_table_name;
```

**修改临时表**

临时表的修改操作与普通表类似，可以使用 ALTER TABLE 命令。

```
ALTER TABLE temp_table_name
ADD COLUMN new_column datatype;
```

**删除临时表**

临时表在会话结束时会自动被销毁，但你也可以使用 DROP TABLE 明确删除它。

```
DROP TEMPORARY TABLE IF EXISTS temp_table_name;
```

#### 二十四、MySQL复制表

如果我们需要完整的复制MySQL的数据表，包括表的结构，索引，默认值等。

如果仅仅使用**CREATE TABLE ... SELECT**命令，是无法实现的。

本章将为介绍如何完整的复制MySQL数据表，步骤如下：

- 使用SHOW CREATE TABLE命令获取创建表(CREATE TABLE)语句，该语句包含了原数据表的结构，索引等。
- 复制以下命令显示SQL语句，修改数据表名，通过以上命令将完全的复制数据表结构。
- 如果你先复制表的内容，你就可以使用**INSERT INTO ... SELECT**语句来实现。

**实例:复制runoob_tbl**

**步骤一：** 获取数据表的完整结构。

```
mysql> SHOW CREATE TABLE runoob_tbl \G;
*************************** 1. row ***************************
       Table: runoob_tbl
Create Table: CREATE TABLE `runoob_tbl` (
  `runoob_id` int(11) NOT NULL auto_increment,
  `runoob_title` varchar(100) NOT NULL default '',
  `runoob_author` varchar(40) NOT NULL default '',
  `submission_date` date default NULL,
  PRIMARY KEY  (`runoob_id`),
  UNIQUE KEY `AUTHOR_INDEX` (`runoob_author`)
) ENGINE=InnoDB 
1 row in set (0.00 sec)

ERROR:
No query specified
```

**步骤二：**

修改 SQL 语句的数据表名，并执行 SQL 语句。

```
mysql> CREATE TABLE `clone_tbl` (
  -> `runoob_id` int(11) NOT NULL auto_increment,
  -> `runoob_title` varchar(100) NOT NULL default '',
  -> `runoob_author` varchar(40) NOT NULL default '',
  -> `submission_date` date default NULL,
  -> PRIMARY KEY  (`runoob_id`),
  -> UNIQUE KEY `AUTHOR_INDEX` (`runoob_author`)
-> ) ENGINE=InnoDB;
Query OK, 0 rows affected (1.80 sec)
```

**步骤三：**

执行完第二步骤后，你将在数据库中创建新的克隆表 clone_tbl。 如果你想拷贝数据表的数据你可以使用 **INSERT INTO... SELECT** 语句来实现。

```
mysql> INSERT INTO clone_tbl (runoob_id,
    ->                        runoob_title,
    ->                        runoob_author,
    ->                        submission_date)
    -> SELECT runoob_id,runoob_title,
    ->        runoob_author,submission_date
    -> FROM runoob_tbl;
Query OK, 3 rows affected (0.07 sec)
Records: 3  Duplicates: 0  Warnings: 0
```

执行以上步骤后，会完整的复制表的内容，包括表结构及表数据。

****

**来给大家区分下mysql复制表的两种方式。**

**第一、只复制表结构到新表**

create table 新表 select * from 旧表 where 1=2

或者

create table 新表 like 旧表 

**第二、复制表结构及数据到新表**

create table新表 select * from 旧表 

****

另一种完整复制表的方法:

```
CREATE TABLE targetTable LIKE sourceTable;
INSERT INTO targetTable SELECT * FROM sourceTable;
```

其他:

可以拷贝一个表中其中的一些字段:

```
CREATE TABLE newadmin AS
(
    SELECT username, password FROM admin
)
```

可以将新建的表的字段改名:

```
CREATE TABLE newadmin AS
(  
    SELECT id, username AS uname, password AS pass FROM admin
)
```

可以拷贝一部分数据:

```
CREATE TABLE newadmin AS
(
    SELECT * FROM admin WHERE LEFT(username,1) = 's'
)
```

可以在创建表的同时定义表中的字段信息:

```
CREATE TABLE newadmin
(
    id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY
)
AS
(
    SELECT * FROM admin
)  
```

#### 二十五、MySQL元数据

MySQL元数据是关于数据库和其对象（如表、列、索引等）的信息。

元数据存储在系统表中，这些表位于MySQL数据库的information_schema数据库中，通过查询这些系统表，你可以获取关于数据库结构、对象和其他相关的详细信息。

你可能想知道MySQL以下三种信息：

- **查询结果信息**：SELECT,UPDATE或DELETE语句影响的记录数。
- **数据库和数据库表的信息**：包含了数据库及数据表的结构信息。
- **MySQL服务器信息**：包含了数据库服务器的当前状态，版本号等。

在MySQL的命令提示符中，我们可以很容易的获取以上服务器信息，单如果使用Perl或PHP等脚本语言，你就需要调用特定的连接函数老获取。

1.查看所有数据库：

```
SHOW DATABASES;
```

2.选择数据库：

```
USE database_name;
```

3.查看数据库中的所有表：

```
SHOW TABLES;
```

4.查看表结构：

```
DESC tabble_name;
```

5.查看表的索引：

```
SHOW INDEX 	FROM table_name;
```

6.查看表的创建语句：

```
SHOW CREATE TABLE table_name;
```

7.查看表的行数：

```
SELECT COUNT(*) FROM table_name;
```

8.查看列的信息：

```
SELECT COLUMN_NAME,DATAE_TYPE,IS_NULLABLE,COLUMN_KEY
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_SCHEMA = 'your_database_name'
ADD TABLE_NAME = 'your_table_name';
```

以上SQL语句中的'your_database_name'和'your_table_name'分别是你的数据库名和表名。

9.查看外键信息：

```
SELECT
	TABLE_NAME,
	COLUMN_NAME,
	CONSTRAINT_NAME,
	REFERENCED_TABLE_NAME,
	REFERENCED_COLUMN_NAME
FROM
	INFORMATION_SCHEMA.KEY_COLUMN_USAGE
WHERE
	TABLE_SCHEMA = 'your_datebase_name'
	ADD TABLE_NAME = 'your_table_name'
	ADD REFERENCED_TABLE_NAME IS NOT NULL;
```

请替换上述SQL语句中的'your_database_name'和'your_table_name'为实际的数据库名和表名。

****

**information_schema数据库**

information_schema是MySQL数据库中的一个系统数据量，它包含有关数据库服务器的元数据信息，这些信息以表的形式存储在information_schema数据库中。

**SCHEMATA表**

存储有关数据库的信息，如数据库名、字符集、排序规则等。

```
SELECT * FROM information_schema.SCHEMATA;
```

### 索引篇

#### 了解索引

##### 索引是什么

**索引的定义：**索引是一个**单独的、存储在磁盘上的数据库**结构（就是一张对应关系的数据表，记录主键和索引），它包含了对数据表所有记录的引用指针。使用索引用于快速找出在某个或多个列中有一特定值的行。对相关列使用索引是提高查询操作速度的最佳途径。

**MySQL索引的提升速度的原理**：索引的建立对于MySQL的高效运行时很重要的，索引可以大大提高MySQL检索速度。比如我们在查字典的时候，前面都有索引的拼音和偏旁、笔画等，然后找到对应字段的页码，这样然后就打开字典的页数就可以知道我们要搜索的某一个key的全部值的信息了。

**合理创建索引**：创建索引时，你需要确保该索引是应用在SQL查询语句的条件(一般作为WHERE子句的条件)，而不是select的字段中，实际上，索引也是一张 ’表‘ ，**该表保存了主键与索引字段**，并指向实体表的记录，虽然索引大大提高了查询速度，同时却会降低更更新表的速度，如对表进行INSERT、UPDATE和DELETE。因为更新表时，MySQL不仅要保存数据，还要保存一下索引文件，建立索引会占用磁盘空间的索引文件。说白了索引就是用来提高速度的，但是就需要维护索引造成资源的浪费，所以合理的创建索引是必要的。

##### 索引的类别

先去官网文档看看支持的索引类型，索引的实现方式如下图所示：https://dev.mysql.com/doc/refman/8.0/en/create-index.html（官网）

![image-20240417095722877](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240417095722877.png)

​		由于是基于MySQL的InnoDB存储引擎，索引类别主要看第一个表格，其他表格可以自行的观看，都不难，从表格我们可以看出来，InnoDB存储引擎索引只支持**BTREE**类型的索引，索引的类别有**Primary Key**, **Unique** , **Key** , **FULLTEXT**和**SPATIAL**。当然也有其他的分法，按照索引列的数量分为**单列索引**和**组合索引**。

1. **Primary KEY(聚集索引)：**InnoDB存储引擎的表会存在一个主键（唯一非null），如果建表的时候没有指定主键，则会使用第一个非空的唯一索引作为聚集索引，否则InnoDB会自动帮你创建一个不可见、长度位6字节的row_id用来作为聚集索引。
2. **单列索引：**单列索引即一个索引只包含单个列
3. **组合索引：**组合索引指在表的多个字段组合上创建索引，只有在查询条件中使用了这些字段的左边字段时，索引才会被使用，使用组合索引时**遵循最左前缀集合**
4. **Unique(唯一索引)：**索引列的值必须唯一，但运行有空值。若是组合索引，则列值的组合必须唯一。主键索引时一种特殊的唯一索引，不允许有空值
5. **Key(普通索引)：**是MySQL中基本索引类型，允许在定义索引的列中插入重复值和空值
6. **FULLTEXT(全文索引)：**全文索引类型为FULLTEXT，在定于索引的列上支持值的全文查找，允许在这些索引列中插入重复值和空值。全文索引可以在CHAR、VARCHAR或者TEXT类型的列上创建
7. **SPATIAL：**空间索引是对空间数据类型的字段建立的索引，MySQL中的空间数据类型有4种，分别是GEOMETRY、POINT、LINESTRING和POLYGON。MySQL使用SPATIAL关键字进行拓展，使能够用于创建正规索引类似的语法创建空间索引。创建空间索引的列必须声明为NOT NULL

这里在说一下组合索引的遵循最左前缀原则：

```
order by使用索引最左前缀
- order by a
- order by a,b
- order by a,b,c
- order by a desc, b desc, c desc 

如果where使用索引的最左前缀定义为常量，则order by能使用索引
- where a=const order by b,c
- where a=const and b=const order by c
- where a=const and b > const order by b,c

不能使用索引进行排序
- order by a , b desc ,c desc  --排序不一致
- where d=const order by b,c   --a丢失
- where a=const order by c     --b丢失
- where a=const order by b,d   --d不是索引的一部分
- where a in(...) order by b,c --a属于范围查询
```

##### 索引的创建原则

因为MySQL本身已经非常优秀了，
在几万条数据的情况之下，索引的优势并不明显。
数据达到几十万条以后，索引的效果显著，能明显提升查询速度，数据量越大，索引越发重要。当数据量有了千万级别时，有无索引可导致性能相差千倍！

1. 索引并非越多越好，一个表种如果有大量的索引，不仅占用磁盘空间，而且会影响INSERT、DELETE、UPDATE等语句性能，因为在表中的数据更改的同时，索引也会进行调整和更新。
2. 避免对经常更新的表进行过多的索引，并且索引中的列尽可能少。而对经常用于查询的字段应该创建索引，但要避免添加不必要字段。
3. 数据量小的表最好不要使用索引，由于数据较少，查询花费的时间可能比遍历索引的时间还要短，索引可能不会产生优化效果。
4. 在条件表达式中，字段的值重复太多的列上不要建立索引。比如在学生表中”性别“字段上只有”男“与”女“两个不同值，因此就无需建立索引。如果建立索引，不但不会提高查询效率，反而会严重减低数据更新速度。
5. 当唯一性是某种数据本身的特征时，指定唯一索引。使用唯一索引需要确保定义的列的数据完整性，以提高查询速度。
6. 在频繁进行排序或分组（即进行group by或order by操作）的列上建立索引，如果待排序的列有多个，可以在这写列上建立组合索引。
7. 搜索的索引列，不一定是所选择的列。换句话说，最适合所有的列是出现在WHERE子句中的列，或链接子句中指定的列，而不是出现在SELECT关键字后的选择列表中的列。
8. 使用短索引。然后对字符串列进行索引，应该指定一个前缀长度，只要有可能就应该这样做。例如，有一个CHAR(200)列，如果在前10个或20个字符内，多数值是唯一的，那么就不要对整个列进行索引。对前10个或20个字符进行索引能够节省大量索引空间，也可能会是查询更快。较小的索引涉及的磁盘IO较少，较短的值比较起来更快。更为重要的是，对于较短的键值，索引高速缓存中的块能容纳更多的键值，因此，MySQL也可以在内存中容纳更多的值。这样就增加了找到行而不用读取索引中较多块的可能性。
9. 利用最左前缀。在创建一个n列的索引时，实际时创建了MySQL可利用的n个索引。多所以可其几个索引的做因，因为可利用多因中最左边的列集来匹配行。这样的列集称为最左前缀。

##### 主键索引和二级索引

**1.聚集索引**

InnoDB存储引擎表时索引组织表，即表中数据按照主键顺序存放。而聚集索引（clustered index）就是按照每张表的主键构建一颗B+树，同时叶子节点中存放的即为整张表的行记录数据，也将聚集索引的叶子节点成为数据页。聚集索引的这个特性决定了索引组织表中数据也是索引的一部分。同B+树数据结构一样，每个数据页都通过了一个双向链表来进行链接。

由于实际的数据页只能按照一颗B+树进行排序，因此每张表只能拥有一个聚集索引。由于定义了数据的逻辑顺序，聚集索引能够特别块地访问指针范围值的查询。查询优化器能够快速发现某一段时间范围的数据页需要扫描。

聚集索引的存储并不是物理上连续的，而是逻辑上连续的。这其中有两点：一是前面说过页通过双向链表链接，页按照主键的顺序排序；另一点是每个页中的记录也是通过双向链表进行维护的，物理存储上可以同样不按照主键存储。

**2.二级索引（辅助索引）**

对于辅助索引（Secondary Index），叶子节点并不包含行记录的全部数据。叶子节点除了包含键值以外，每个叶子节点中的索引中还包含一个书签（bookmark）。该书签用来高速InnoDB存储引擎哪里可以找到索引相对应的行数据的聚集索引键。

当通过辅助索引来寻找数据时，InnoDB存储引擎会遍历辅助索引并通过叶级别的指针获得指向主键索引的主键，然后再通过主键索引来找一个完整的行记录。

**3.覆盖索引**

InnoDB存储引擎支持覆盖索引（conering index，或称索引覆盖），即从辅助索引中就可以得到查询的记录，而不需要查询聚集索引的记录。使用覆盖索引的好处是辅助索引不包含整行记录的所有信息，故其大小要原小于聚集索引，因此可以减少大量的IO操作。

##### 索引实现的原理

InnoDB存储引擎的索引是基于B+树实现的，从索引类别表格可以看出，不支持hash的实现方式。首先来了解B+树的特点。

**B+树的特征：**

1. 有K个子树的中间节点包含K个元素（B树中式K-1个元素），每个元素不保存数据，只用来索引，所有数据都保存在叶子节点。
2. 所有的叶子节点中包含了全部元素的信息，及指向含这些元素记录的指针，且叶子节点本身依关键字的大小自小而大顺序链接。
3. 所有的中间节点元素都同时存在与子节点，在子节点元素中是最大（最小）元素。

**B+树的优势：**

1. 单一节点存储更多的元素，使得查询的IO次数更少。
2. 所有查询都要查找到叶子节点，查询性能稳定。
3. 所有叶子节点形成有序链表，便于范围查询。

在B+树中，所有记录节点都是按键值的大小顺序存放在同一层叶子节点上，由各叶子节点指针进行连接。    先来看一个B+树，其中高度为2，每页可存放4条记录，扇出（fan out）为5，如下图所示：

![image-20240417145353689](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240417145353689.png)

**索引的设计思考：**

- 索引是一种存储方式，最相关的硬件就是磁盘，索引磁盘的性能会直接影响到数据库查询的效率
- 磁盘的性能的读写的顺序有关，普通磁盘顺序读写比随机读写快很多，所以尽量避免随机读写。
- 数据都是以行为单位一行一行的存储的，每一行都包括了所有的列，多行可以连续存储。
- 每一行数据中，一般都有一个键，其他的列可以称为值，可以理解为键值对。InnoDB必须有唯一非空的主键，就是默认的键。
- 在键值对中，键值可以排序，还可以组合键值。

**索引的设计：**

- 磁盘空间会划分许多各大小相等的块或页，一个页中可以存储多行数据，这样就符合磁盘的顺序读写，这样一次IO就可以读取很多数据到内存，就可以减少磁盘IO。
- 在一个页内，所有的数据可能会经常变动，并且大小也是相对固定的，所以内部通过链表或数组管理。
- 每个键值可以排序，所以在一个块内的所有数据也是可以是有序的，这样通过二分查找法查找可以很多的在一个页内找到指定键对应的数据。
- 一个页设计好之后，可以把页作为B+树的节点，通过页来承载数据，通过B+树来组织不同页之间的关系。
- B+树的特点是在内节点存储键来提高搜索的性能，所以很自然的，内节点用来存储数据行的键，叶子节点存储作用数据行，可以很好的提升性能。

表中数据按照主键顺序存放。而聚集索引（clustered index）就是按照每张表的主键构造一棵B+树，同时叶子节点中存放的即为整张表的行记录数据，也将聚集索引的叶子节点称为数据页。聚集索引的这个特性决定了索引组织表中数据也是索引的一部分。同B+树数据结构一样，每个数据页都通过一个双向链表来进行链接。如下图所示：

![image-20240417155514869](C:\Users\FWR\AppData\Roaming\Typora\typora-user-images\image-20240417155514869.png)

​		上图所示的是一个深度为2的B+树，也是我们所称的索引，这里假设页有随机唯一的编号，根页号为20。这里有一个内节点（根节点），其他的都是叶子节点，也是数据节点，对于内节点来说，存有key和pageno的指针信息，对于叶子节点来说，只存有完整的数据。对于聚集索引，data部分存有除主键外的其他列的组合，如果是二级索引，则这里存放就是这行记录对应主键的组合，用于**回表。**

​		最左边的**MIN**为了很好的组织树形结构的指针，和其他内的节点一样，主要用来标记它是最小记录**Min**，还有就是个pageno指针指向下层最左边的Min记录，其他的节点的Min记录用于判断搜索是否到了边界。每个页都有页头页尾和标记页面的状态，页面中的数据是如何存储，有没有空闲的空间，以什么样的顺序存储等。

​		上图中所有的叶子节点从左到右都是从小到大顺序以双向链表的方式存储的，所以当我们需要遍历全部的数据，之需要通过B+树找到最小的位置，然后通过遍历链表则可以查询到所有的数据，还有就是10，16，25这三条记录在内节点和叶子节点均存在，这即使B+树的特点，叶子节点会存有所有的**key和值**。而内节点只存储了**key**，不存储其他的数据，只有用来索引。叶子节点除了第一条记录会有上一层重复的存储，其他数据不会有这样的现象，所以浪费的空间页不大，由于每个页的大小都是固定的（16K），在内节点上只存储key，不存储其他数据，一个页就可以存储更多的key，这样检索也能减少磁盘的IO，由于页存储key增多，这样就可以使得B+树的深度减少，这样也可以减少磁盘的IO，提高查询性能。

例如一个三层的B+树，每一页都能存1000个key，所以第二层就有1000*（1+1000）个key，第三层就有1000 * 1001 *1001 = 1002001000（十亿级别），一个简单的三层B+数据就可以存放十亿级别的数据，很强大。

​		上面说到的**回表**其实就是在使用二级索引进行搜索时，因为二级索引只保存了部分列的数据，如果需要获取键值不包括的列的数据时，需要通过二级索引的指针（书签：用于指向聚集索引的指针）找到聚集索引的全部数据，然后返回需要查询的列的值。如果使用二级索引不能找到需要的值（需要回表），成为非覆盖索引，否则为覆盖索引。非覆盖索引需要回表，增加IO，所以性能会差一些。所以可以根据业务需求创建组合索引来避免回表。但是也要权衡索引带来的利是否大于弊。所以在统计行总数的时候可以通过二级索引来统计，这样速度回块一些。大概图形如下：

![image-20240417165534892](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240417165534892.png)

怎么去理解覆盖索引和非覆盖索引对应的二级索引关系呢？

- **覆盖索引：**例如一张学生表，根据学生的学号查出姓名，学号和姓名都作为二级索引（辅助索引），这种情况下，在辅助索引就可以查到姓名了，就不用去聚集索引进行回表，根据书签（叶级别的指针指向主键索引）来找一个完整的行记录。、
- **非覆盖索引**：检索完辅助索引没有得到对应列的数据，根据辅助索引得到的书签去聚集索引查找得到完整行的记录，这种操作称为回表

这里所一些不能走索引的情况，但是不多说，因为优化这个东西太多，后其在深入了解，虽然可能创建了很多索引，很多情况都不走索引，比如：like %query_name%、where端使用or条件连接，where端使用函数等，在group by和order by使用的时候注意组合索引的最左前缀原则。

#### 为什么采用B+树作为索引

##### 与hash表和b树的比较

B+树、哈希表和B树是三种常用的数据结构，用于数据库索引、文件系统以及内存中的数据管理。每种结构各有其优势和特定的使用场景。

**B+树**

B+树是B树的一个变体，广泛用于数据库和文件系统中，主要因为其高效的磁盘存取性能和对范围查询的支持。

**优点**：

1. **所有值都在叶子节点**：这使得范围查询变得非常高效，因为可以通过叶子节点的链接直接遍历这些值。
2. **更少的磁盘访问**：由于其宽度大，深度通常较浅，从而减少了磁盘I/O操作。
3. **顺序访问优化**：叶节点之间是相互链接的，这对顺序访问非常有利。

**缺点**：

- 相对复杂的插入和删除操作，因为这些操作可能需要节点的分裂或合并。

****

**哈希表**

哈希表提供了非常快速的点查找能力，通常用于支持快速查找操作的场景。

**优点**：

1. **快速查找性能**：理想情况下，哈希表的查找、插入和删除操作的时间复杂度为O(1)。
2. **直接关键字访问**：不需要遍历数据结构就能直接找到元素。

**缺点**：

- **不支持范围查询**：哈希表不适合进行范围搜索或有序数据访问。
- **哈希冲突**：当多个键映射到同一个哈希值时，会发生冲突，这需要额外的解决机制。
- **动态扩展问题**：当哈希表填满时，需要重新哈希所有元素到一个更大的表中，这是一个高成本的操作。

****

**B树**

B树是一种自平衡树，以其多路平衡搜索树的特性被广泛应用于数据库和操作系统。

**优点**：

1. **自平衡**：B树通过在插入和删除时进行节点分裂和合并，保持平衡状态。
2. **有效的磁盘访问**：B树设计用来工作在存储系统中，有效减少磁盘I/O，节点通常大小与磁盘块相等。
3. **支持顺序和范围查询**：通过中序遍历，B树可以有效地进行范围查询。

**缺点**：

- 比B+树有更多的指针跳转，因为数据可以存储在内部节点，导致范围查询可能不如B+树高效。

****

**总结**

- **B+树**通常用于数据库索引，特别是当需要高效支持范围查询时。
- **哈希表**适用于需要快速查找的场景，但不支持顺序或范围查找。
- **B树**在需要自平衡和支持范围查询的系统中使用，但其操作复杂性和磁盘访问次数多于B+树。

选择哪种数据结构取决于应用的具体需求，包括数据访问模式、预期的查询类型以及性能要求。

#### 索引失效情况有哪些

是有filesort的情况（说白了就是不走索引）：

1. where语句与order by语句，使用了不同的索引
2. 检查的行数过多，且没有使用覆盖索引
3. ORDER BY中的列不包含在相同的索引，也就是使用了不同的索引
4. 对索引列同时使用了ASC和DESC
5. where语句或者ORDER BY语句中索引列使用了表达式，包括函数表达式
6. where 语句与ORDER BY语句组合满足最左前缀，但where语句中使用了条件查询。

#### 优化索引的方法

**一、优化能够带来什么**

优化能带来的肯定是爽！！！当你作为一个用户去体验某个网页正在加载中。。。内心肯定是崩溃的，秒出的网页肯定是爽的。做开发测试的每次查询都是几分钟甚至十几分钟，内心不会崩溃，只想砸电脑，所以速度快就是爽。上面提到的优化可以从多个方面进行：最常见的就是sql和索引的优化，因为写CRUD的小伙伴避免不了写一个查询语句，然后语句走不走索引呢，这也会决定你爽不爽。

**二、优化思路**

既然要做优化，首先要知道哪些需要优化吧，**首先定位慢sql语句，**然后做分析这个sql慢是因为什么，然后才是怎么解决呢，在实际的环境中可能先要定位慢的语句，然后观察一段时间是不是一直都慢呢，还是有时候慢有时候快，这些都和实际环境中的并法、Mysql环境当时的健康程度有关，索引要先锁定目标然后观察，然后把哪些慢的sql都弄出来进行分析，最后做优化。

**慢查询语句（mysqldumpslow）**

mysql中有很多日志文件，binlog日志，慢日志，查询日志，错误日志。这里我们要说的是慢日志，默认情况下是没有开启慢日志查询日志的，因为会影响一些性能。

#### explain关键字怎么分析，看哪些字段

**explain作用：**

1. 可以知道表的执行顺序
2. 在sql中哪些索引可以使用
3. 在sql中哪些索引实际上被用
4. 数据读取操作的操作类型
5. sql中的每个表有多少行数据被优化器查询
6. 表之间的引用

**explain怎么玩：**

```
语法：
explain + sql查询

---
explain select * from item_description td inner join (select * from item_general where item_id in (select item_id from item where item_id > 332604504321036693 and item_id < 332604504321036710)) tt on tt.item_id=td.item_id;
```

![image-20240418153831550](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240418153831550.png)

**下面看explain的具体说明：**

1. **id：**select查询的序号列，表示查询select语句中表的执行顺序

   - 当id相同，则从上往下执行
   - 如果id不同，则从大到小顺序执行
   - 如果id有相同，则从大到小的顺序执行

2. **select_type**：表示SELECT语句的类型。它可以是以下几种取值：

   - SIMPLE：表示简单查询，其中不包括连接查询和子查询；
   - PRIMARY：表示主查询，或者是最外层的查询语句，最外层查询为PRMARY，也就是最后加载的就是PRIMARY；
   - UNION：表示来连接查询的第二个或后面的查询语句，不依赖外部查询的结果集
   - UNION RESULT:连接查询的结果；
   - SUBQUERY:子查询中的第1个SELECT语句；不依赖于外部查询的结果集
   - DEPENDENT SUBQUERY:子查询中的第1个SELECT，依赖于外面的查询；
   - DERIVED:导出表的SELECT（FROM子句的子查询）,MySQL会递归执行这些子查询，把结果放在临时表里。
   - DEPENDENT DERIVED:派生表依赖于另一个表
   - MATERIALIZED:物化子查询
   - UNCACHEABLE SUBQUERY:子查询，其结果无法缓存，必须针对外部查询的每一行重新进行评估
   - UNCACHEABLE UNION:UNION中的第二个或随后的 select 查询，属于不可缓存的子查询

   3、**table**:表示查询的表

   4、**partitions**:查询将从中匹配记录的分区。该值适用NULL于未分区的表

   5、**type**:表示表的连接类型

   - system:该表是仅有一行的系统表。这是const连接类型的一个特例
   - const: 数据表最多只有一个匹配行，它将在查询开始时被读取，并在余下的查询优化中作为常量对待。const表查询速度很快，因为只读取一次,const用于使用常数值比较PRIMARY KEY或UNIQUE索引的所有部分的场合。
   - eq_ref:对于每个来自前面的表的行组合，从该表中读取一行,可以用于使用=运算符进行比较的索引列 。比较值可以是常量，也可以是使用在此表之前读取的表中列的表达式
   - ref:对于来自前面的表的任意行组合，将从该表中读取所有匹配的行，ref可以用于使用“＝”或“＜＝＞”操作符的带索引的列。
   - fulltext:使用FULLTEXT 索引执行联接
   - ref_or_null:这种连接类型类似于ref，但是除了MySQL还会额外搜索包含NULL值的行。此联接类型优化最常用于解析子查询
   - index_merge:此联接类型指示使用索引合并优化。在这种情况下，key输出行中的列包含使用的索引列表，并key_len包含使用的索引 的最长键部分的列表
   - unique_subquery:类型替换 以下形式的eq_ref某些 IN子查询,unique_subquery 只是一个索引查找函数，它完全替代了子查询以提高效率。
   - index_subquery:连接类型类似于 unique_subquery。它代替IN子查询,但只适合子查询中的非唯一索引
   - range:只检索给定范围的行，使用一个索引来选择行。key列显示使用了哪个索引。key_len包含所使用索引的最长关键元素。当使用＝、＜＞、＞、＞＝、＜、＜＝、IS NULL、＜＝＞、BETWEEN或者IN操作符用常量比较关键字列时，类型为range
   - index:该index联接类型是一样的 ALL，只是索引树被扫描。这发生两种方式：1、如果索引是查询的覆盖索引，并且可用于满足表中所需的所有数据，则仅扫描索引树。在这种情况下，Extra列显示为 Using index，2、使用对索引的读取执行全表扫描，以按索引顺序查找数据行。 Uses index没有出现在 Extra列中。
   - ALL:对于前面的表的任意行组合进行完整的表扫描
   - system>const>eq_ref>ref>index>ALL(这些比较常见，)

   6、**possible_keys**:指出MySQL能使用哪个索引在该表中找到行。若该列是NULL，则没有相关的索引。在这种情况下，可以通过检查WHERE子句看它是否引用某些列或适合索引的列来提高查询性能。如果是这样，可以创建适合的索引来提高查询的性能。

   7、**kye**:表示查询实际使用的索引，如果没有选择索引，该列的值是NULL。要想强制MySQL使用或忽视possible_keys列中的索引，在查询中使用FORCE INDEX、USE INDEX或者IGNORE INDEX

   8、**key_len**：表示MySQL选择的索引字段按字节计算的长度，若键是NULL，则长度为NULL。注意，通过key_len值可以确定MySQL将实际使用一个多列索引中的几个字段

   9、**ref**：表示使用哪个列或常数与索引一起来查询记录。

   10、**rows**：显示MySQL在表中进行查询时必须检查的行数。

   11、**filtered**：按表条件筛选的行的百分比

   12、**Extra**：表示MySQL在处理查询时的详细信息

   - using filesort: 使用了文件排序，很影响性能
   - using temporary: 使用了临时表，很影响性能
   - using index: 使用了覆盖索引
   - using where: 使用了where
   - using MRR: 使用了MRR优化
   - using join buffer: 使用了链接缓存
   - impossible where: where的条件总是false
   - select tbles optimized away: 在没有group by操作的情况下，不必等到执行阶段再计算
   - distnct: 优化distinct操作
   - using index condition: 使用了ICP优化

****

#### count(*)和count(1)有何区别？哪个性能更高？

在MySQL中，`COUNT(*)`和`COUNT(1)`通常用来计算表中的行数，但它们在语义上略有不同，尽管在大多数情况下，它们的性能几乎是相同的。

**语义区别**

- **`COUNT(*)`**：这个语句会计算表中的所有行，不考虑列的值。`COUNT(*)`被用来统计包含NULL值的行，因为它简单地计数所有行。
- **`COUNT(1)`**：这个语句也是用来计数表中的所有行，但是数字1在这里仅作为一个常数出现。这意味着它不是基于任何特定列的值进行计数，也包括所有行，就像`COUNT(*)`一样。在这里使用1，或任何非NULL常数（如`COUNT(0)`，`COUNT(-1)`），都会得到相同的结果。

**性能比较**

对于`COUNT(*)`和`COUNT(1)`的性能而言，现代的数据库管理系统（包括MySQL）在优化这类查询时都非常智能。在执行计划中，这两种形式几乎没有区别，因为数据库优化器会识别出这两个表达式的实际用途——仅计数行，并生成相同的执行计划。

在早期的数据库系统中，有一些讨论认为`COUNT(1)`可能会比`COUNT(*)`稍快一些，因为`*`可能需要检查所有列，而`1`是一个简单的常量。然而，这在现代数据库系统中已不再是问题，因为优化器足够智能，可以理解这两个查询实际上是相同的。

**实践建议**

- **使用`COUNT(\*)`**：由于`COUNT(*)`是SQL标准中明确用于计数行的语法，因此建议使用`COUNT(*)`来表示你的意图更加清晰。这使得代码更容易理解，尤其是对于那些熟悉SQL标准的人来说。
- **测试和优化**：如果你在特定情况下怀疑性能问题，可以实际测试`COUNT(*)`与`COUNT(1)`在你的具体环境和数据库版本中的性能。但在大多数情况下，你可能不会看到任何显著差异。

总之，从实用和标准的角度来看，`COUNT(*)`是首选方法，而从性能角度来看，两者之间没有显著差异。



### 事务篇

**事务的概念**：事务（transaction）是一个操作序列，这些操作要么全部执行，要么全部不执行。

事务通常以BEGIN TRANSACTION开始，以COMMIT或ROLLBACK操作结束。

**事务的特性：**

1. **原子性（atomicity）**：指事务在逻辑上是不可分隔的操作单元，其所有语句要么都执行，要么都撤销执行。
   举例：
   假设又两个账户，A账户和B账户。A账户转给B账户100元，这里有两个动作在里面：A账户减去100元，B账户增加100元，这两个动作不可分隔，即原子性。
2. **一致性（consistency）：**事务时一种逻辑上的工作单元。一个事务就是一系列在逻辑上相关的操作指令的集合，用于完成一项任务，其本质就是将数据库的数据从一种一致性状态换到另一种一致性状态。
   举例：
   假设A账户和B账户两者的钱加一起一共是5000元，那不管A账户和B账户之间如何转账，转了几次帐，事务结束后，两个用户的钱加起来应该还得是5000元，这就是事务的一致性。
3. **隔离性（isolation）：**隔离是针对并发事务而言的，所谓并发是指数据库服务器同时处理多个事务，如果不采取专门的控制机制，那么并发之间可能会相互干扰，进而导致数据出现不一致或错误的状态。
   隔离性就是隔离并发运行的多个事务间的相互影响。关于事务的隔离性，数据库提供了多种隔离级别，各有优缺点，下面会展开分析
   **隔离性要达到的效果：**对于任意两个并发的事务T1、T2，在事务T1看来，T2要么在T1开始之前就结束，要么在T1结束后才开始。这样每个事务都感觉不到有其他事务在并发执行。
4. **持久性：**指一旦事务提交成功，其对数据的修改就是持久性的。
   数据更新的结果已经从内存转到外部存储器上，此后即使发生了系统故障，已提交的事务所作的数据更新也不会丢失。

**为什么要隔离，如果不考虑事务的隔离，可能会发生什么？**

1. **脏读（dirty read）：**
   一个事务读取了已被另一个事务修改，但尚未提交的数据。

   当一个事务正在多次修改某个数据，而这个事务这多次的修改都还没提交，这时，另一个并发的事务来访问数据时，就会造成两个事务得到的数据不一致
   举例：
   update account set money = money + 1000 where name = 'B';

   update account set money = money - 1000 where name = 'A';

   当只执行成功了第一条SQL时，整个事务还未执行成功，也没有提交事务，这时另一个并发事务来查询B账户的钱，得到的结果就是B账户增加了1000元之后的结果。这个时候就形成了脏读，无论第二条SQL是否执行，只要该事务不提交，那么所有操作都将回滚，这时也就意味着之前另一个并发事务查询的结果是不正确的。

2. **不可重复读（nonrepeatable read）：**
   在同一事务中，同一个查询在Time1时刻读取某一行，在Time2时刻重新读取这一行的数据时，发现这个行的数据已经发生修改，可能被更新了（UPDATE），也可能被删除了（DELETE）。
   举例：
   事务T1在读取某一行数据，而事务T2立即修改了这个数据，并提交事务给数据库，事务T1再次读取该数据就得到了不同的结果，发生了不可重复读。

3. **幻读（phantom read**）：
   在同一事务中，当同一查询多次执行的时候，由于其他插入（insert）操作的事务提交，会导致每次返回不同的结果集。
   举例：
   事务T1对一个表中的所有行的某个字段执行了从1修改到2的操作，这是事务T2又在这个表中插入了一行数据，而这行数据的那个字段还是1，并提交给数据库。
   这时，操作事务T1的用户如果在查看刚刚修改的数据，就会发现还有一条没有修改，就好像产生幻觉一样，这就是发生了幻读（其实这行数据是另一个事务T2新添加的）

**事务的四种隔离级别**

1. **未提交读（read uncommitted）**
   在该隔离级别，所有事务都可以看到其他未提交事务的处理结果，即在未提交读级别，事务中的修改，即使没有提交，对其他事务也是可见的。
   该隔离级别很少用于实际应用，隔离级别最低，并发性最高，但会出现脏读，不可重复读、幻读三种问题。
2. **提交读（read committed）**
   它满足了隔离的基本定义：一个事务只能看见已经提交事务所做的改变。
   换句话说，一个事务从开始到提交之前，所做的任务修改对其他事务都是不可见的。
   它是Oracle、SQL Server数据库系统的默认隔离级别。不会出现脏读，但会出现不可重复读和幻读问题。
3. **可重复读（repeatable read）**
   可重复读可以确保同一个事务，在多次读取同样的数据时，得到同样的结果。
   可重复读是MySQL数据库系统的默认隔离级别。它不会出现脏读和不可重复读的问题，但仍会出现幻读问题。
   MySQL数据库中的InnoDB和FaIcon存储引擎通过**MVCC（multi-version concurrent control 多版本并发控制）**解决了不可重复读问题，而MVCC加上间隙锁又解决了幻读问题。
4. **可串行化、序列化（serializable）**
   这是最高的隔离级别，它通过强制事务排序，强制事务串行执行。使之不可能相互冲突。
   它不会出现脏读、不可重复读和幻读问题。但是可能会出现大量的超时现象和锁竞争。
   实际应用中很少用到这个级别。只有在非常需要数据的一致性，且接受无并发的情况下，才考虑用该级别。

****

#### MVCC是什么

MVCC，全称Multi-Version Concurrent Control，即多版本并发控制。MVCC是一种并发控制的方法，一般在数据库管理系统中，实现对数据库的并发访问，在编程语言中实现事务内存。

MVCC在MySQL InnoDB中的实现主要是为了提高数据库并发性能，用更好的方式去处理读-写冲突，做到即使有读写冲突时，也能做到不加锁，非阻塞并发读。

学习并发版本控制之前，了解当前读和快照读

- **当前读**
  像select lock in share mode（共享锁），select for update；update，insert，delete（排他锁）这些操作都是一种当前读，为什么叫当前读？就是它读取的记录的最新版本，读取时还要保证其他并发事务不能修改当前记录，会对读取的记录进行加锁
- **快照读**
  像不加锁的select操作就是快照读，即不加锁的非阻塞读；快照读的前提是隔离级别不是串行级别，串行级别的快照读会退化成当前读；之所以会出现快照读的情况，是基于提高并发性能的考虑，快照读的实现是基于多版本并发控制，即MVCC，可以认为MVCC是行锁的一个变种，但它在很多情况下，避免了加锁操作，减低了开销；既然是基于多版本，即快照读可能读到的并不一定是数据的最新版本，而有可能是之前的历史版本。

> 说白了MVCC就是为了实现读-写冲突不加锁，而这个读指的就是快照读，而非当前读，当前读实际是一种加锁的操作，是悲观锁的实现

- 准确的说，MVCC多版本并发控制指的是"维持一个数据的多个版本，使得读写操作没有冲突"这么一个概念。仅仅是一个理想概念
- 而在MySQL中，实现这么一个MVCC理想概念，我们就需要提供具体的功能去实现它，而快照读就是MySQL为我们实现MVCC理想模型的其中一个具体非阻塞读功能，当前读就是悲观锁的具体功能实现

**MVCC能解决什么问题。好处是？**

数据库并发场景可以分为三种：

1. **读-读：**不存在任何问题，也不需要并发控制
2. **读-写：**有线程安全问题，可能会造成事务隔离性问题，可能遇到脏读，幻读，不可重复读
3. **写-写：**有线程安全问题，可能会存在更新丢失问题，比如第一类更新丢失，第二类更新丢失

- readview是什么
- rr和rc隔离级别下，readview的生成时机
- undolog版本链怎么生成的
- readview和undolog怎么比较的——四个判断依据
- mcvv能完全解决幻读问题吗

#### 可重复读隔离级别，完全解决了幻读吗

### 锁篇

#### MySQL有哪些锁

![image-20240429103938601](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240429103938601.png)

##### 全局锁

- **全局锁的应用场景是什么？**

全局锁主要应用于做全库逻辑备份，这样在备份数据库期间，不会因为数据或表结构的更新，而出现备份文件的数据于预期的不一样。

举个例子大家就知道了。

在全库逻辑备份期间，假设不急啊全局锁的场景，看看会出现什么意外的情况。

如果在全库备份期间，有用户购买了意见商品，一般购买商品的业务逻辑会涉及到多张数据库表的更新，比如在用户表更新该用户的余额，然后在商品表被购买的商品的库存。

那么，有可能出现这样的顺序：

1. 先备份了用户表的数据；
2. 然后有用户发起了购买商品的操作；
3. 接着在备份商品表的数据。

也就是在备份用户表和商品表之间，有用户购买了商品。

这中情况下，备份的结果是用户表该用户的余额并没有扣除，反而商品表中该商品的库存被减少了，如果后面用这个备份文件恢复数据库数据的话，用户钱没少，而库存少了，等于用户白嫖了一件商品。

所以，在全库逻辑备份期间，加上全局锁，就不会出现上面这张情况了。

- **全局锁是怎么用的？**

要使用全局锁，则要执行这条命令：

```
flush tables with read lock
```

执行后，**整个数据库就处于只读状态了**，这时其他线程执行以下操作，都会被阻塞：

1. 对数据的增删改操作，比如insert、delete、update等语句；
2. 对表结构的更改操作，比如alter table、draop table等语句。

如果要释放全局锁，则要执行这条命令：

```
unlock tables
```

当然，当会话断开了，全局锁会被自动释放

在异常处理机制上有差异，如执行 FTWRL 命令之后由于客户端发生异常断开，那么 MySQL 会自动释放这个全局锁，整个库回到可以正常更新的状态。而将整个库设置为 readonly 之后，如果客户端发生异常，则数据库就会一直保持 readonly 状态，这样会导致整个库长时间处于不可写状态，风险较高。

记住，业务的更新不只是增删改数据（DML），还有可能是加字段等修改表结构的操作（DDL）。不论是哪种方法，一个库被全局锁上之后，你要对里面任何一个表做更改操作，都会被锁住的。

- **加上全局锁会带来哪些缺点呢？**

加上全局锁，意味着整个数据库都是只读状态。

那么如果数据库里面有很多数据，备份就会花费很多的时间，关键是备份期间，业务只能读取数据，而不能更新数据，这样会造成业务停滞。

- **既然备份数据库数据的时候，使用全局锁会影响业务，那有什么其他方式可以避免？**

有的，如果数据库的引擎支持的事务支持可重复读的隔离级别，那么在备份数据库之前先开启事务，会先创建Read View，然后整个事务执行期间都在用这个Read View，而且由于MCVV的支持，备份期间业务依然可以对数据进行更新操作。

因为在可重复读的隔离级别下，即使其他事务更新了表的数据，也不会影响备份数据库时的Read View，这就是事务的四大特性中的隔离性，这样备份期间备份的数据一直在开启事务时的数据。

备份数据库的工具是mysqldump，在使用mysqldump时加上`-single-tansaction`参数的时候，就会在备份数据库之前先开启事务。这种方法只适用于支持[可重复读隔离级别的事务]的存储引擎。

InnoDB存储引擎默认的事务隔离级别正是可重复读，因此可以采用这种方式来备份数据库。

但是，对于MyISAM这种不支持事务的引擎，在备份数据库时就要使用全局锁的方法。

*****

##### 表级锁

> MySQL表级锁有哪些？具体怎么用的。

MySQL里面表级别的锁有这几种：

- 表锁；
- 元数据锁（MDL）；
- 意向锁；
- AUTO-INC锁；

###### **表锁**

先来说说**表锁**。

如果我们想对学生表（t_student）加表锁，可以使用下面的命令：

```
//表级别的共享锁，也就是读锁；
lock tables t_student read;

//表级别的独占锁，也就是写锁；
lock tables t_student write;
```

需要注意的是，表锁除了会限制别的线程的读写外，也会限制本线程接下来的读写操作。

也就是说如果本线程对学生表加了[共享表锁]，那么本线程接下来如果要对学生表执行写的操作的语句，是会被阻塞的，当然其他线程对学生表进行写操作时也会被阻塞，直到锁被释放。

要释放表锁，可以使用下面这条命令，会释放当前会话的所有表锁：

```
unlock tables
```

另外，当会话退出后，也会释放所有表锁。

不过尽量避免在使用InnoDB引擎的表使用表锁，因为表锁的颗粒度太大，会影响并发性能，**InnoDB牛逼的地方在于实现了颗粒度更细的行级锁**。

###### **元数据锁（MDL）**

我们不需要显示的使用MDL，因为当我们对数据库进行操作时，会自动给这个表加上MDL：

- 对一张表进行CRUD操作时，加的是MDL读锁；
- 对一张表做结构变更操作的时候，加的是MDL写锁；

MDL是为了保证当前用户对表执行CRUD操作时，防止其他线程对这个表结构做了变更。

当有线程执行select语句（加MDL读锁）的期间，如果有其他线程要更改该表的结构（申请MDL写锁），那么将会被阻塞，直到执行完select语句（释放MDL读锁）。

反之，当有线程对表结构进行变更（加MDL写锁）的期间，如果有其他线程执行了CRUD操作（申请MDL读锁），那么就会阻塞，直到表结构变更完成（释放MDL写锁）。

- **MDL不需要显示调用，那它是在什么时候释放的？**

MDL是在事务提交后才会释放，这意味着**事务执行期间，MDL是一直持有的**。

那如果数据库有一个长事务（所谓长事务，就是开启了事务，但是一直没提交），那在对表结构做变更操作的时候，可能会发生意想不到的事情，比如下面这个顺序的场景：

1. 首相，线程A先启用了事务（但是一直不提交），然后执行一条select语句，此时就先对该表加上了MDL读锁；
2. 然后，线程B也执行了同样的select语句，此时并不会阻塞，因为[读读]并不冲突；
3. 接着，线程C修改了表字段，此时由于线程A的事务并没有提交，也就是MDL读锁还占用着，这时线程C就无法申请到MDL写锁，就会被阻塞，

那么在线程C阻塞后，后续有对该表的select语句，就都会被阻塞，如果此时有大量该表的select语句的请求到来，就会有大量的线程被阻塞住，这时数据库的线程很快就会爆满了。

- **为什么线程C因为申请不到MDL写锁，而导致后续的申请读锁的查询操作也会被阻塞？**

这是因为申请MDL锁的操作会形成一个队列，队列中**写锁获取优先级高于读锁，**一旦出现MDL写锁等待，会阻塞后续该表的所有CRUD操作。

所有为了能安全的对表结构进行变更，在对表结构变更前，先要看看数据库的长事务，是否事务已经对表加上了MDL读锁，如果可以考虑kill掉这个长事务，然后做表结构的变更。

###### **意向锁**

1. 在使用InnoDB引擎的表里对某些记录加上[共享锁]之前，需要先在表级别加上一个[意向共享锁]；
2. 在使用InnoDB引擎的表里对某些记录加上[独占锁]之前，需要先在表级别加上一个[意向独占锁]；

也就是，当执行插入、更新、删除操作，需要先对表加上[意向独占锁]，然后对该记录加独占锁。而普通的select是不会加行级别锁的，普通的select语句是利用MVCC实现一致性读，是无锁的。

不过，select也是可以对记录加共享锁和独占锁的，具体方式如下：

```
//先在表加上意向共享锁，然后对读取的记录加共享锁
select ... lock in share mode;

//先表上加上意向独占锁，然后对读取的记录加独占锁
select ... for update;
```

意向共享锁和意向独占锁是表级别，不会和行级的共享锁和独占锁发生冲突，而且意向锁之间也不会发生冲突，只会和共享表锁（lock tables ... read）和独占表锁（lock tables ... write）发生冲突。

表锁和行锁是满足读读共享、读写互斥、写写互斥的。

那么有了[意向锁]，由于在对记录加独占锁前，先回加上表级别的意向独占锁，那么在加[独占表锁]，直接查该表是否有意向独占锁，如果有就意味着表里已经有记录被加了独占锁，这样就不用去遍历表里的记录。

所以，意向锁的目的是为了快速判断表里是否有记录被加锁。 

###### **AUTO-INC锁**

表里的主键通常都会设置成自增的，这是通过对主键字段声明`AUTO_INCREMENT`属性实现的。

之后可以在插入数据时，可以不指定主键的值，数据库回自动给主键赋值递增的值，这主要时通过AUTO-INC锁实现的。

AUTO-INC锁是特殊的表的表锁机制，**锁不是在一个事务提交后才释放，而是在执行完插入语句后就会立即释放**。

**在插入数据时，会加一个表级别的AUTO-INC锁**，然后为被**AUTO_INCREMENT**修饰的字段赋值递增的值，等插入语句执行完成后，才会把AUTO-INC锁释放掉。

那么，一个事务在持有AUTO-INC锁的过程中，其他事务的如果要向该表插入语句都会被阻塞，从而保证插入数据时，被AUTO_INCREMENT修饰的字段的值是连续递增的。

但是，AUTO-INC锁再对大量数据进行插入的时候，会影响插入性能，因为另一个事务中的插入会被阻塞。

因此，再MySQL 5.1.22版本开始，InnoDB存储引擎提供了一种轻量级的锁来实现自增。

****

##### 行级锁

InnoDB引擎是支持行级锁的，而MyISAM一七年并不支持行级锁。

行级锁的类型主要有三类：

- Record Lock，记录锁，也就是仅仅把一条记录锁上；
- Gap Lock，间隙锁，锁定一个范围，但是不包含记录本身；
- Next-Key Lock：Record Lock + Gap Lock的组合，锁定一个i范围，并且锁定记录本身。

###### Record Lock

Record Lock称为记录锁，锁住的是一条记录。而且记录锁是有S锁和X锁之分的：

- 当一个事务对一条记录加了S型记录锁后，其他事务也可以继续对该记录加S型记录锁（S型与S锁兼容），但是不可以对该记录加X型记录锁（S型与X锁不兼容）；
- 当一个事务对一条记录加了X型记录锁后，其他事务既不可以对该记录加S型记录锁（S型与X锁不兼容），也不可以对该记录加X型记录锁（X型与X锁不兼容）。
- 举个例子，当一个事务执行了下面这条语句：

```
select * from t_test where id = 1 for update;
```

就是对t_test表中主键id为1的这条记录加上X型的记录锁，这样其他事务就无法对这条记录进行修改了。

![image-20240429172322931](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240429172322931.png)

###### Gap Lock

Gap Lock称为间隙锁，只存在于可重复读隔离级别，目的是为了解决可重复读隔离级别下幻读的现象。

假设，表中有一个范围id为（3，5）间隙锁，那么其他事务就无法插入id = 4这条记录了，这样就有效的防止幻读现象的发生。

![image-20240429172647274](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240429172647274.png)

间隙锁索然存在X型间隙锁和S型间隙锁，但是并没有声明区别，**间隙锁之间是兼容的，即两个事务可以同时持有包含通过间隙范围的间隙锁，并不存在互斥关系，因为间隙锁的目的是防止插入幻影记录而提出的。**

###### Next-Key Lock

Next-Key Lock称为临键锁，是Record Lock + Gap Lock的组合，锁定一个范围，并且锁定记录本身。假设表中有一个范围id为3(3,5]的next-key lock，那么其他事务既不能插入id - 4记录，也不能够修改id = 5这条记录。

![image-20240429173210782](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240429173210782.png)

所以，next-key lock既能保护该记录，又能阻止其他事务将新纪录插入到被保护记录前面的间隙中。

**next-key lock是包含间隙锁 + 记录锁如果一个事务获取了X型的next-key lock，那么另外一个事务在获取相同范围的X型的next-key lock时，是会被阻塞的。**

比如，一个事务持有了范围为(1,10]的X型的next-key lock，那么另外一个事务在获取相同范围的X型的next-key lock时，就会被阻塞。

虽然相同范围的间隙是多个事务相互兼容的，但对于记录锁，我们是需要考虑X型于S型关系，X型的记录锁与X型的记录锁是冲突的。

#### MySQL是怎么加行级锁的？

行级锁加锁规则比较复杂，不同的场景，加锁的形式是不同的。

**加锁的对象是索引，加锁的基本单位是next-key lock，它是由记录锁和间隙锁组合而成的，next-key lock是前后闭区间，而间隙锁是前后开区间。**

但是，next-key lock在一些场景下会退化成记录锁或间隙锁。

那到底是什么场景呢？总结一句，**在能使用记录锁或者间隙锁就能避免幻读现象的场景下，next-key lock就会退化成记录锁或间隙锁。**

一下面这个表结构进行试验说明：

```mysql
CREATE TABLE `user(
	`id` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
    `age` int NOT NULL,
    PRIMATY KEY (`id`),
    KEY `index_age` (`age`) USTING BTREE
)ENGINE=InnoDB DUFAULT CHARSET-utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

其中，id是主键索引（唯一索引），age是普通索引（非唯一索引），name是普通的列。

表中的有这些行记录：

![image-20240430152853173](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240430152853173.png)

不同版本的加锁规则可能是不同的，但是大体上是相同的。

##### 唯一索引等值查询

当我们用唯一索引进行等值查询的时候，查询的记录不存在，加锁的规则也会不同：

- 当查询的记录是[存在]的，在索引树上定位打牌这一条记录后，将该记录的索引中的next-key lock会**退化成[记录锁]**
- 当查询的记录是[不存在]的，在索引树找到第一条大于该查询记录的记录后，将该记录的索引中的nex-key lock会**退化成[间隙锁]。**

> tip
>
> 本篇文章的[唯一索引]是用[主键索引]作为案例说明的，加锁只加在主键索引项上。
>
> 然后，很多同学误以为如果是二级索引的[唯一索引]，加锁也是只加二级索引项上。
>
> 其实是不对的，这里特别说明下，如果是用二级索引(不管是不是非唯一索引，还是唯一索引)进行锁读查询的时候，除了会对二级索引项加行级锁(如果是唯一索引的二级索引，加锁规则和主键索引的案例相同），而且还会对查询到的记录的主键索引项加[记录锁]。
>
> 在文章的[非唯一索引]的案例中，使用二级索引作为例子，在后面的章节有说明，对二级索引进行锁定读查询的时候，因为存在两个索引（二级索引和主键索引），所以两个索引都会加锁。



#### UPDATE没加索引会锁全表？

#### MySQL记录锁+间隙锁 可以防止删除操作而导致的幻读吗

#### MySQL死锁了怎么办

#### 加了什么锁，导致死锁的

#### MySQL怎么查看表锁的情况

### 日志篇

Mysql日志：undo log、redo log、binlog有什么用？

下面查询语句经历的过程，这属于[读]一条记录的过程，如下图：

![image-20240521094509402](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240521094509402.png)

那么，执行一条update语句，期间发生了什么？比如下面这一条update语句：

```
UODATE t_user SET name = 'xiaomi' WHERE id = 1;
```

查询语句的那一套流程，更新语句也是同样会走一遍：

- 客户端先通过连接器建立连接，连接器自会判断用户身份；
- 因为这是一条update语句，所以不需要经过查询缓存，但是表上有更新语句，是会把整个表的查询缓存清空的，所以说查询缓存很鸡肋，在Mysql8.0就被一吃这个功能了；
- 解析器会通过语法分析识别出关键字update，表名等等，构建出语法树，接着还会做语法分析，判断输入的语句是否符合MySQL语法；
- 预处理器会判断表和字段是否存在；
- 优化器确定执行计划，因为where条件中的id是主键索引，所以决定要使用id这个索引；
- 执行器负责具体执行，找到这一行，然后更新。

不过更新语句流程会涉及到undo log（回滚日志）、redo log（重做日志）、binlog（归档日志）这三种日志：

- **undo log（回滚日志）：**是Innodb存储引擎层生成的日志，实现了事务中的**原子性**，主要用于**事务回滚和MVCC**；
- **redo log（重做日志）：**是Innodb存储引擎层生成的日志，实现了**事务中的持久性**，主要用于**掉电等故障恢复**；
- **binlog（归档日志）：**是server层生成的日志，主要用于数据备份和主从复制；

所以这次带着这些问题，看看这三种日志怎么工作的。

![image-20240521102437843](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240521102437843.png)

#### 为什么需要undo log（回滚日志）？

我们在执行一条"增删改"语句的时候，虽然没有输入begin开启事务和commit提交事务，但是MySQL会**隐式开启事务**来执行"增删改"语句的，执行完成自动提交事务的，这样就保证了执行完"增删改"语句后，我们可以及时在数据库表看到"增删改"的结果了。

执行一条语句是否自动提交事务，是由**autocommit**参数决定的，默认是开启。所以，执行一条update语句也会使用事务的。

那么，考虑一个问题。一个事务在执行过程中，在还没有提交事务之前，如果MySQL发生了崩溃，要怎么回滚事务之前的数据呢？

如果我们每次在事务执行过程中，都记录下回滚时需要的信息到一个日志里，那么在事务执行中途发生了MySQL崩溃后，就不用担心无法回滚到事务之前的数据，我们可以通过这个日志回滚到事务之前的数据。

实现这一机制就是**undo log（回滚日志），它保证了事务的ACID特性中的原子性(Atomicity)**。

undo log是一种用于撤销回退的日志。在事务还没提交之前，MySQL会先记录更新前的数据到undo log日志文件里面，当事务回滚时，可以利用undo log来进行回滚，如下图：

![image-20240521104619734](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240521104619734.png)

每当InnoDB引擎对一条记录进行操作（修改、删除、新增）时，要把回滚时需要的信息都记录到undo log里，比如：

- 在**插入**一条记录时，要把这条记录的主键记录下来，这样之后回滚时只需要把这个主键值对应的记录删掉就好了；
- 在**删除**一条记录时；要把这条记录的内容都记下来。这样之后回滚时再把这些内容组成的记录**插入**到表中就好了；
- 在**更新**一条记录时，要把被更新的列的旧值记下来，这样之后回滚时再把这些列**更新为旧值**就好了。

在发生回滚时，就读取undo log里的数据，然后做原先相反操作。比如当delete一条记录时，undo log中会把记录中的内容都记录下来，然后执行回滚操作的时候，就读取undo log里的数据，然后进行insert操作。

不同的操作，需要记录的内推也是不同的，所以不同类型的操作（修改、删除、新增）产生的undo log的格式也是不同的，具体的每一个操作的undo log的格式就不详细介绍了。

一条记录的每一次更新操作产生的undo log格式都有一个roll_pointer指针和一个trx_id事务id：

- 通过trx_id可以知道该记录时被那个事务修改的；
- 通过roll_pointer指针可以将这些undo log串成一个链表，这个链表就别成为版本链；

版本链 如下图：

![image-20240521111721818](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240521111721818.png)

另外，undo log**还有一个作用，通过readView + undo log实现MVCC（多版本并发控制）**。

对于[读提交]和[可重复读]隔离级别的事务来说，它们的快照读（普通select语句）是通过Read View + undo log 来实现的，它们的区别在于创建Read View的时机不同：

- [读提交]隔离级别是在每个select都会产生一个新的Read View，也意味着，事务期间的多次读取同一条数据，前后两次读取的数据可能会出现不一致，以为可能这期间另外一个事务修改了该记录，并提交了事务；
- [可重复读]隔离级别是启动事务时生成一个Read View，然后整个事务期间都在用这个Read View，这样就保证了早事务期间读到的数据都是事务开启前的记录。

这两个隔离界别实现是通过[事务的Read View里的字段]和[记录中的两个隐藏列（trx_id和roll_pointer）]的比对，如果不满足可见行，就会顺着undo log版本链里找到满足其可见性的记录，从而控制并发事务访问同一个记录时的行为，这就叫**MVCC**(多版本并发控制)。

因此，undo log作用：

- **实现事务回滚，保障事务的原子性**，事务处理过程中，如果出现了错误或者用户执行了ROOLLBACK语句，MySQL可以利用undo log中的历史数据恢复到事务开始之前的状态。
- 实现**MVCC（多版本并发控制）关键因素之一**。MVCC是通过Read View + undo log实现的。undo log为了每条记录保存多分历史数据，MySQL在执行快照读（普通select语句）的时候，会根据事务的Read View里的信息，顺着undo log的版本链找到满足其可见性的记录。

> TIP
>
> 很多人疑问undo log是如何刷盘（持久化到磁盘）的？
>
> undo log和树叶的刷盘策略是一样的，都需要通过redo  log保证持久化。
>
> buffer pool中有undo页，对undo页的修改也都会记录到redo log。redo  log会每秒刷盘，提交事务时也会刷盘，数据页和undo页都是靠这个机制保证持久化的。

#### 为什么需要Buffer Pool？

****

MySQL的数据库都是存在磁盘中的，那么我们要更新一条记录的时候，得先要从磁盘读取该记录，然后在内存中修改这条记录，那修改完这条记录是悬着直接写回到磁盘，还是选择缓存起来呢？

当然是缓存起来好，这样下次有查询语句命中了这条记录，直接读取缓存中的记录，就不不用从磁盘读取数据了。

为此，InnoDB存储引擎设计了一个**缓冲池（Buffer Pool）**，来提高数据库的读写性能。

![image-20240521143739424](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240521143739424.png)

有了Buffer Pool后：

- 当读取数据时，如果数据存在于Buffer Pool中，客户端就会直接读取Buffer Pool中的数据，否则再去读磁盘中读取。
- 当修改数据时，如果数据存于Buffer Pool中，那直接修改Buffer Pool中数据所在的页，然后将其页设置为脏页（改页的内存数据和磁盘的数据已经不一致），为了减少磁盘I/O，不会立即将脏页写入磁盘，后续由后台线程选择一个合适的时机将脏页写入到磁盘。

Buffer Pool缓存什么？

InnoDB会把存储的数据划分为若干个[页]，以页作为磁盘和内存交互的基本单位，一个页默认大小为16KB。因此，Buffer Pool同样需要按[页]来划分。

在Mysql启动的时候，**InnoDB会为Buffer Pool申请一篇连续的内存空间，然后按照默认 的16KB的大小划分出一个个的页，Buffer Pool中的页就叫做缓存页。**此时这些缓存页都是空闲的，之后随着程序的运行，才会有磁盘上的页被缓存到Buffer Pool中。

所以，MySQL刚启动的时候，你会观察到使用的虚拟内存空间很大，而世道的物理内存空间却很小，这是因为只有这些虚拟内存被访问后，操作系统才会触发缺页中断，申请物理内存，接着将虚拟地址和物理地址建立映射关系。

Buffer Pool除了缓存[索引页]和[数据页]，还包括了Undo页，插入缓存、自适应哈希索引、锁信息等。

![image-20240521154501549](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240521154501549.png)

> undo页是记录什么？

开启事务后，InnoDB层更新记录前，首先要记录相应的undo log，如果是要更新操作，需要把被更新的列的旧值记下来，也就是生成一条undo log，undo log会写入Buffer Pool中的undo页面。

> 查询一条记录，就需要缓冲一条记录嘛？

不是的。

当我们查询一条记录时，InnoDB是会把整个页的数据加载到Buffer Pool中，将页加载到Buffer Pool后，在通过页里的[页目录]去定位到某条具体的记录，

为什么需要redo log（重做日志）？

Buffer Pool是提高了读写的效率没错，但是问题来了，Buffer Pool是基于内存的，而内存总是不可靠的，万一断电重启，还没来得及落盘的脏页数据就会丢失。

为了防止断电导致数据丢失的问题，当有一条记录需要更新的时候，InnoDB引擎就会先更新内存（同时标记为脏页），然后将本次对这个页的修改以redo log的形似记录下来，**这个时候更新就算完成了。**

后续，InnoDB引擎会在适当的时候，由后台将缓存在Buffer Pool的脏页刷新到磁盘里，这就是WAL（Writer-Ahead Logging）技术。

**WAL技术指的是，MySQL的写操作并不是立刻写到磁盘上的，而是先写日志，然后在合适的时间再写到磁盘上。**

过程如下图：

![image-20240521160005494](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240521160005494.png)

> 什么是redo log?

redo log是物理日志，记录了某个数据页做了什么修改，比如**对xxx表空间中的yyy数据页zzz偏移量的地方做了AAA更新**，每当执行一个事务就会产生这样的一条或多条物理日志。

在事务提交时，只要先将redo log持久化到磁盘即可，可以不需要等到将缓存在Buffer Pool里的脏页数据持久化到磁盘。

当系统崩溃时，虽然脏页数据没有持久化，但是redo log已经持久化，接着MySQL重启后们，看一看根据redo log的内容，将所有的数据恢复到最新的状态。

> 被修改Undo页面，需要记录对应的redo log吗？

需要的。

开启事务后，InnoDb层更新记录前，首先要记录相应的undo log，如果是更新操作，需要把被更新的列的旧值记下来，也就是要生成一条undo log，undo log会写入Buffer Pool中的Undo页面。

不过，**在内存修改该Undo页面后，需要记录对应的redo log。**

> redo log和undo log区别在哪？

这两种日志是属于InnoDB存储引擎的日志，它们的区别在于：

- redo log记录了此次事务[完成后]的数据状态，记录的是**更新之后的值**；
- undo log记录了此次事务[开始前]的数据状态，记录的是**更新之前的值**；

事务提交之前发生了崩溃，重启后会通过undo log回滚事务，事务提交之后发生了崩溃，重启后会通过redo log恢复事务，如下图：

![image-20240521161733101](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240521161733101.png)

所以有了redo log，再通过WAL技术，InnoDB就可以保证及时数据库发生异常重启，之前以及调的记录都不会丢失，这个能力成为crash-safe（崩溃恢复）。可以看出来，**redo log保证了事务四大特性中的持久性。**

> redo log要写到磁盘，数据也要写到磁盘，为什么要多此一举?

写入redo log的方式使用了追加操作，所以磁盘操作时顺序写，而写入数据需要先找到写入位置，才能写到磁盘，所以磁盘操作是**随机写。**

磁盘的[顺序写]比[随机写]高效的多，因此redo log写入磁盘的开销更小。

针对[顺序写]为什么比[随机写]更快这个问题，可以比喻为你有一个本子，按照顺序一页一页写肯定比写一个字都要找到对应页写的快得多。

可以说这是WAL技术的另外一个优点：MySQL的写操作从磁盘的[随机写]变成了[顺序写]，提升语句的执行性能。这是因为MySQL的写操作并不是立刻更新到磁盘上，而是先记录在日志上，然后再核实的时间再更新到磁盘上。

至此，针对为什么需要redo log这个问题我们有两个答案：

- 实现事务的持久性，让MySQL有crash-safe的能力，能够保证MySQL在任何时间段突崩溃，重启后之前已提交的记录都不会丢失；
- 将写的操作从[随机写]变成了[顺序写]，提升MySQL写入磁盘的性能。

> 产生的redo log是直接写入磁盘的吗？

不是的。

实际上，执行一个事务的过程中，产生的redo log也不是直接写入磁盘的，因为这样会产生大量的I/O操作，而且磁盘的运行速度远慢于内存。

所以，redo log也有自己的缓存——redo log Buffer，每当产生一条redo log时，会先写到redo log buffer，后续在持久化到磁盘，如下图：

![image-20240521173538045](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240521173538045.png)

redo log buffer 默认大小16MB，可以通过innodb_log_Buffer_size参数动态的调整大小，增大它的大小可以让Mysql处理[大事务]是不必写入磁盘，进而提升写IO性能。

##### redolog什么时候刷盘？

缓存在redo log buffer里的redo log还是在内存中，它什么时候刷新到磁盘？

主要有下面几个时机：

- MySQL正常关闭时；
- 当redo log buffer中记录的写入量大于redo log buffer内次空间的一半时，会触发落盘；
- InnoDB的后台线程每隔1s，将redo log buffer持久化到磁盘。
- 每次事务提交时都将缓存在redo log buffer里的redo log直接持久化到磁盘（这个策略可由innodb_flush_log_at_trx_commit参数控制）

>innodb_flush_log_at_trx_commit参数控制的是什么？

单独执行一个更新语句的时候，InnoDB引擎会自启动一个事务，在执行更新语句的过程中，生成的redo log先写入到redo log buffer中，，然后等事务提交的时候，再讲缓存在redo log buffer中的redo log buffer中的redo log按组的方式[顺序写]到磁盘。

上面这种redo log刷盘时机是在事务提交的时候，这个默认的行为。

除此之外，InnoDB还提供了另外两种策略，由参数

`innodb_flush_log_at_trx_commit` 参数控制，可取的值有：0、1、2，默认值为 1，这三个值分别代表的策略如下：

- 当设置该**参数为 0 时**，表示每次事务提交时 ，还是**将 redo log 留在 redo log buffer 中** ，该模式下在事务提交时不会主动触发写入磁盘的操作。
- 当设置该**参数为 1 时**，表示每次事务提交时，都**将缓存在 redo log buffer 里的 redo log 直接持久化到磁盘**，这样可以保证 MySQL 异常重启之后数据不会丢失。
- 当设置该**参数为 2 时**，表示每次事务提交时，都只是缓存在 redo log buffer 里的 redo log **写到 redo log 文件，注意写入到「 redo log 文件」并不意味着写入到了磁盘**，因为操作系统的文件系统中有个 Page Cache（如果你想了解 Page Cache，可以看[这篇 (opens new window)](https://xiaolincoding.com/os/6_file_system/pagecache.html)），Page Cache 是专门用来缓存文件数据的，所以写入「 redo log文件」意味着写入到了操作系统的文件缓存。

##### 为什么需要binlog？

前面介绍了undo log 和 redo log这两个日志都是InnoDB存储引擎生成的。

MySQL在完成一条更新操作后，Server层会生成一条binlog，等之后事务提交的时候，会将事务执行工程中产生的所有binlog同意写入binlog文件。

binlog文件是记录了所有数据结构变更和表数据修改的日志，不会记录查询类的操作，比如SELECT和SHOW操作。

> 为什么有了binlog，还要redo log？

这个问题跟MySQL的时间线有关系。

最开始MySQL里并没有InnoDB引起，MySQL自导的引擎是MylSAM，但是MylSAM没有crash-safe的能力，binlog日志只能用于归档。

而InnoDB是另一个公司以插件形式引入MySQL的，既然只依靠binlog是没有crash-safe能力的，所以InnoDB使用了redo log来实现crash-safe能力。



##### redo log 与binlog有什么区别？

这两个日志有四个区别：

1. 适用对象不同
   - binlog是MySQL的server层实现的日志，所有存储引擎都可以使用；
   - redo log是InnoDB存储引擎实现的日志；
2. 文件格式不同：
   - redo log 是物理日志，记录的是在某个数据页做了什么修改，比如对 XXX 表空间中的 YYY 数据页 ZZZ 偏移量的地方做了AAA 更新；
   - binlog 有 3 种格式类型，分别是 STATEMENT（默认格式）、ROW、 MIXED
3. 写入方式不同：
   - binlog是追加写，写满一个文件，就创建一个新的文件继续写，不会覆盖以前的日志，保存的是全量的日志。
   - redo log是循环写，日志空间大小是固定的，全部写满就从头开始，保存未被刷入磁盘的脏页日志。
4. 用途不同：
   - binlog用于备份恢复、主从复制；
   - redo log用于掉电等故障恢复。

> 如果不小心整个数据库的数据被删除了，能使用redo log文件恢复数据吗？

不可以使用redo log文件恢复，只能使用binlog文件恢复。

因为redo log文件是循环写，是会边写边擦除日志的，之记录未被输入磁盘的数据的物理日志，已经刷入磁盘的数据都会从redo     log文件擦除。

binlog文件保存的是全量的日志，也就是保存了所有数据变更的情况，理论上只要记录在binlog上的数据，都可以恢复，所以如果不小心整个数据库的数据被删除了，得用binlog文件恢复数据。

#### 主从复制是怎么实现？

MySQL的主从复制以来binlog，也就是记录MySQL上的所有变化并以二进制保存在磁盘上。复制的过程就是将binlog中的数据从主库传输到从库上。

这个过程一般是异步的，也就是主库上执行事务操作的线程不会等待复制binlog的线程同步完成。

![image-20240523144519415](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240523144519415.png)

**MySQL集群的主从复制过程梳理成3个阶段：**

- 写入binlog：主库写binlog日志，提交事务，并更新本地存储数据。
- 同步binlog：把binlog复制到所以从库上，每个从库把binlog写到暂存日志中。
- 回放binlog：回放binlog，并更新存储引擎中的数据。

**具体纤细过程如下：**

- MySQL主库在收到客户端提交事务的请求之后，会先写入binlog，再提交事务，更新存储引擎中的数据，事务提交完成后，返回给客户端“操作成功”的响应。
- 从库会创建一个专门的I/O线程，连接主库的log dump线程，来接收主库的binlog日志，再把binlog信息写入relay log的中继日志里，再返回给主库"复制成功"的响应。
- 从库会创建一个用于回放binlog的线程，去读relay log中继日志，然后回放binlog更新存储引擎中的数据，最终实现主从的数据一致性。

完成主从复制之后，你就可以在写数据时只写主库，在读数据时只读从库，这样及时写请求会锁表或者锁记录，也不会影响读请求的执行。

![image-20240523145906564](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240523145906564.png)

> 从库还不是越多越好？

不是的。

因为从库数量增加，从库连接上来的I/O线程页比较多，**主库也要创建同样多的log dump线程来处理复制的请求，对主库资源小号比较高，同时还受限于主库的网络带宽。**

所以在实践使用中，一个主库一般跟2~3个从库(1套数据库，1主2从1备主)，这就是一主多从的MySQL集群架构。

> MySQL主从复制还有哪些模型？

主要有三种：

- **同步复制**：MySQL 主库提交事务的线程要等待所有从库的复制成功响应，才返回客户端结果。这种方式在实际项目中，基本上没法用，原因有两个：一是性能很差，因为要复制到所有节点才返回响应；二是可用性也很差，主库和所有从库任何一个数据库出问题，都会影响业务。
- **异步复制**（默认模型）：MySQL 主库提交事务的线程并不会等待 binlog 同步到各从库，就返回客户端结果。这种模式一旦主库宕机，数据就会发生丢失。
- **半同步复制**：MySQL 5.7 版本之后增加的一种复制方式，介于两者之间，事务线程不用等待所有的从库复制成功响应，只要一部分复制成功响应回来就行，比如一主二从的集群，只要数据成功复制到任意一个从库上，主库的事务线程就可以返回给客户端。这种**半同步复制的方式，兼顾了异步复制和同步复制的优点，即使出现主库宕机，至少还有一个从库有最新的数据，不存在数据丢失的风险**。

##### binlog什么时候刷盘？

事务执行过程中，先把事务日志写到binlog cache(Server层的cache)，事务提交的时候，再把binlog cache写到binlog文件中。

一个事务的binlog是不能被拆开的，因此无论这个事务有多大(比如有很多条语句)，也要保证一次性写入。这是因为有一个线程只能同时有一个事务在执行的设定，所以每当执行一个begin/start transaction的时候，就会默认提交上一个事务，这人如果一个事务的binlog被拆开的时候，在备库执行就会被很多个事务分段执行，这样破坏了原子性，是有问题的。

MySQL给每个线程分配了一片内存用于缓存binlog，该内存叫binlog cache，参数binlog_cache_size用于控制单个线程内binlog cache所占内存的大小。如果超过了这个参数规定的大小，就要暂存到磁盘。

> 什么时候binlog cache会写到binlog文件？

在事务提交的时候，执行器把binlog cache里的完整事务写入到binlog文件中，并情况binlog cache。如下图：

![image-20240523170204445](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240523170204445.png)

虽然每个线程有自己的binlog cache，但是最终都写到同一个binlog文件：

- 图中的 write，指的就是指把日志写入到 binlog 文件，但是并没有把数据持久化到磁盘，因为数据还缓存在文件系统的 page cache 里，write 的写入速度还是比较快的，因为不涉及磁盘 I/O。
- 图中的 fsync，才是将数据持久化到磁盘的操作，这里就会涉及磁盘 I/O，所以频繁的 fsync 会导致磁盘的 I/O 升高。

MySQL提供一个 sync_binlog 参数来控制数据库的 binlog 刷到磁盘上的频率：

- sync_binlog = 0 的时候，表示每次提交事务都只 write，不 fsync，后续交由操作系统决定何时将数据持久化到磁盘；
- sync_binlog = 1 的时候，表示每次提交事务都会 write，然后马上执行 fsync；
- sync_binlog =N(N>1) 的时候，表示每次提交事务都 write，但累积 N 个事务后才 fsync。

> 三个日志讲完了，至此我们可以先小结下，update语句的执行过程。

当优化器分析出成本最小的执行的计划后，执行器就按照执行计划开始进行更新操作。

具体更新一条记录`UPDATE t_user SET name = 'xiaoming' WHERE id = 1;`的流程如下：

1.  执行器复制具体执行，会调用存储引起的接口，通过主键索引树搜索获取id = 1这一行记录：
   - 如果id = 1这一行所在的数据页本来就在buffer pool中，就直接返回给执行器更新；
   - 如果记录不再buffer pool，将数据页从磁盘读入到buffer pool，返回记录给执行器。
2. 执行器得到聚簇索引记录后，会看一下更新前的记录和更新后的记录是否一样：
   - 如果一样的话就不进行后续更新流程；
   - 如果不一样的话就把更新前的记录和更新后的记录都当做参数传给InnoDB层，让InnoDB真正执行更新记录的操作；
3. 开启事务，InnoDB层更新记录前，首先要记录相应的undo log，因为这是更新操作，需要把更新的列的旧值记下来，也就是要生成一条undo log，undo log会写入buffer pool中的Undo页面，不过在内存修改Undo页面后，需要记录对应的redo log。
4. InnoDB层开始更新记录，会先更新内存（同时标记为脏页），然后将记录写到redo log里面，这个时候跟新就算完成了。为了减少磁盘I/O，不会立即将脏页写入磁盘，后续有后台线程选择一个合适的时机将脏页写入到磁盘。这就是WAL技术，MySQL的写操作并不是立刻写到磁盘上，而是先写redo日志，然后在合适的时间再将修改的行数据写到磁盘上。
5. 至此，一条记录更新完了。
6. 在一条更新语句执行完成后，然后开始记录该语句对应的binlog，此时记录的binlog会被保存到binlog cache，并没有刷新到硬盘上的binlog文件，在事务提交时才会统一将该事务运行过程中的所有binlog刷新到硬盘。
7. 事务提交，剩下就是[两阶段提交]的事情了。

##### 为什么需要两阶段提交？

事务提交后，redo log和binlog都要持久化到磁盘，但是这两个是独立的逻辑，可能出现半成功的状态，这样就造成两份日志之间的逻辑不一致。

举个例子，假设id =  1这行数据的字段name的值原本是'jay'，然后执行`update t_user set name = 'xiaoming' where id = 1;`如果在持久化redo log和binlog两个日志的过程中，出现了办成功状态，那么就有两种情况：

- **如果在将rredo log刷入到磁盘之后，MySQL突然宕机了，而binlog还没有来得及写入。**MySQL重启后，通过redo log能将buffer pool中的id = 1这行数据的name字段恢复到xiaoming，但是binlog里面没有记录这条更新语句，在主从架构中binlog会被复制到从库，由于binlog丢失了这条更新语句，从库的这一行name字段是旧值jay，与主库的值不一致性；
- **如果在将binlog输入磁盘后，MySQL突然宕机了，而redo log还没有来得及写入。**由于redo log还没写，崩溃恢复以后这个事务无效，所以id = 1这行数据的name还是旧值jay，而binlog里面记录了这条新的语句，在主从架构中，binlog会被复制到从库，从库执行力 这条更新语句，那么这一行name字段是新值xiaoming，与主库的值不一致；

可以看到，在持久化redo log和binlog这两份日志的时候，如果出现了半成功状态，就会造成主从环境的数据不一致性。这是因为redo log影响主库数据，binlog影响从库的数据，所以redo log和binlog必须保持一致才能保证主从数据一致。

**MySQL为了避免出现两份日志之间的逻辑不一致的问题，使用了[两阶段提交]来解决**，两阶段提交其实是分布式事务一致性协议，它可以保证多个逻辑操作要不全部成功，要不全部失败，不会出现半成功的状态。

**两个阶段提交把单个事务的提交拆分成了2个阶段，分别是[准备（Prepare）阶段]和[提交Commit阶段]**，每个阶段都有协调者(Coordinate)和参与者(Participant)共同完成。注意，不要把提交(Commit)阶段和commit语句混淆了，commit语句执行的时候，会包含提交阶段(Commit)阶段。

举个拳击比赛的例子，两位拳手（参与者）开始比赛之前，裁判(协调者）会在中间确认两位拳击手的状态，类似于问你准备好了吗？

- **准备阶段：**裁判(协调者)会依次询问两位拳击手(参与者)首付准备好了，然后拳击手听到后做出应答，如果觉得自己准备好了，就会跟裁判说准备好了；如果自己还没有准备好(比如全套还没有带好)，就会跟裁判说还没准备好。
- **提交阶段：**如果两位拳击手(参与者)都回答准备好了，裁判(协调者）宣布比绍正式开始，两位拳击手直接开打；如果任何异味拳击手(参赛者)回答没有准备好，裁判(协调者)会宣布比赛暂停，对应事务中的回滚操作。

##### 两阶段提交的过程是怎样的？

在MySQL的InnoDB存储引擎中，开启binlog的情况下，MySQL会同时维护binlog日志与InnoDB的redo log，为了保证这两个日志的一致性，MySQL使用了**内部XA事务**（是的，也有外部XA事务，跟本文不太相关，就不介绍了），内部XA事务有binlog作为协调者，存储引擎是参与者。

当客户端commit语言或者在自动提交的情况下，MySQL内部开启了一个XA事务，分阶段来完成XA事务的提交，如下图：

![image-20240526111918418](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526111918418.png)

从图中可看出，事务的提交有两个阶段，**就是将redo log的写入拆成了两个步骤：prepare和commit，中间在穿插写入binlog，**具体如下：

- **prepare阶段：**将XID（内部事务的ID）写入到redo log，同时将redo log对应的事务状态设置为prepare，然后将redo log持久化到磁盘（innodb_flush_log_at_trx_commit =1的作用）；
- **commit阶段：**把XID写入到binlog，然后将binlog持久化到磁盘(sync_binglog = 1的作用)，接着调用引擎的提交事务接口，将redo log状态设置问commit，此时该状态并不需要持久化到磁盘，只需要Write搭配文件系统的page cache中就够了，因为只要binlog写入磁盘成功，就算redo log的状态还是prepare也没有关系，一样会被认为事务已经执行成功；

##### 异常重启会出现什么现象？

我们来看看在两阶段提交的不同时刻，MySQL异常重启会出现什么现象？下图中有时刻A和时刻B都有可能发生崩溃：

![image-20240526112807568](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526112807568.png)

不管是时刻A（redo log已经写入磁盘，binlog还没写入磁盘），还是时刻B（redo log和binlog都已经写入磁盘，还没写入commit标识）崩溃，**此时的redo log都处于prepare状态。**

在MySQL重启会按照顺序扫描redo log文件，碰到处于prepare状态的redo log，就拿着redo log中的XID去binlog查看是否存在此XID：

- 如果binlog中没有当前内部XA事务的XID，说明redo log完成刷盘，但是binlog还没有刷盘，则回滚事务。对应时刻A崩溃恢复的情况。
- 如果binlog中有当前内部XA事务的XID，说明redo log和binlog都已经完成了刷盘，则提交事务。对应的时刻B崩溃恢复的情况。

可以看到，对于处于prepare阶段的redo log，既可以提交事务，也可以回滚事务，这取决于是否能在binlog中查找到于redo log相同的XID，如果有就提交事务，如果没有就回滚事务。这样就可以保证redo log和binlog这两份日志的一致性。

所以说，两阶段提交是以binlog写成功为事务提交成功的标识，因为binlog写成功了，就意味着能在binlog中查找到redo log相同的XID。

> 处于prepare阶段的redo log加上完整的binlog，重启就提交事务，MySQL为什么要这么设计？

binlog已经写入了，之后就会被从库(或者用这个binlog恢复出来的库)使用。

所以，在主库上也要提交这个事务。采用这个策略，主库和备库的数据就保证了一致性。

> 事务没提交的时候，redo log会被持久化到磁盘吗？

会的。

事务执行中间过程的redo log也是直接写在redo log buffer中的，这些缓存的redo log buffer里的redo log也会被[后台线程]每隔一起持久化到磁盘。

也就是说，**事务没提交的时候，redo log也是可能被持久化到磁盘的。**

有的同学可能会问，如果MySQL崩溃了，还没提交事务的redo log已经被持久化磁盘了，MySQL重启后，数据不就不一致了？

放心，这种情况MySQL重启会进行回滚操作了，因为事务没提交的时候，binlog是还没持久化到磁盘的。

##### 两阶段提交有什么问题？

两阶段提交虽然保证了两个日志文件的数据一致性，但是性能很差，主要有两个方面的影响：

- 磁盘I/O次数高：对于“双1”配置，每个事务提交都会进行两次fsync（刷盘），一次是redo log刷盘，另一次是binlog刷盘。
- 锁竞争激烈：两阶段提交虽然能够保证[单事务]两个日志的内容一致，但是[多事务]的情况下，却不能宝泽会给你两者的提交顺序一致，因此，在两阶段提交的流程基础上，还需要加一个锁来保证提交的原子性，从保证多事务的情况下，两个日志的提交顺序一致。

> 为什么两阶段提交的磁盘 I/O 次数会很高？

binlog 和 redo log 在内存中都对应的缓存空间，binlog 会缓存在 binlog cache，redo log 会缓存在 redo log buffer，它们持久化到磁盘的时机分别由下面这两个参数控制。一般我们为了避免日志丢失的风险，会将这两个参数设置为 1：

- 当 sync_binlog = 1 的时候，表示每次提交事务都会将 binlog cache 里的 binlog 直接持久到磁盘；
- 当 innodb_flush_log_at_trx_commit = 1 时，表示每次事务提交时，都将缓存在 redo log buffer 里的 redo log 直接持久化到磁盘；

可以看到，如果 sync_binlog 和 当 innodb_flush_log_at_trx_commit 都设置为 1，那么在每个事务提交过程中， 都会**至少调用 2 次刷盘操作**，一次是 redo log 刷盘，一次是 binlog 落盘，所以这会成为性能瓶颈。

> 为什么锁竞争激烈？

在早期的 MySQL 版本中，通过使用 prepare_commit_mutex 锁来保证事务提交的顺序，在一个事务获取到锁时才能进入 prepare 阶段，一直到 commit 阶段结束才能释放锁，下个事务才可以继续进行 prepare 操作。

通过加锁虽然完美地解决了顺序一致性的问题，但在并发量较大的时候，就会导致对锁的争用，性能不佳。

###### 组提交组提交

**MySQL 引入了 binlog 组提交（group commit）机制，当有多个事务提交的时候，会将多个 binlog 刷盘操作合并成一个，从而减少磁盘 I/O 的次数**，如果说 10 个事务依次排队刷盘的时间成本是 10，那么将这 10 个事务一次性一起刷盘的时间成本则近似于 1。

引入了组提交机制后，prepare 阶段不变，只针对 commit 阶段，将 commit 阶段拆分为三个过程：

- **flush 阶段**：多个事务按进入的顺序将 binlog 从 cache 写入文件（不刷盘）；
- **sync 阶段**：对 binlog 文件做 fsync 操作（多个事务的 binlog 合并一次刷盘）；
- **commit 阶段**：各个事务按顺序做 InnoDB commit 操作；

上面的**每个阶段都有一个队列**，每个阶段有锁进行保护，因此保证了事务写入的顺序，第一个进入队列的事务会成为 leader，leader领导所在队列的所有事务，全权负责整队的操作，完成后通知队内其他事务操作结束。

![image-20240526134054488](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526134054488.png)

对每个阶段引入了队列后，锁就只针对每个队列进行保护，不再锁住提交事务的整个过程，可以看的出来，**锁粒度减小了，这样就使得多个阶段可以并发执行，从而提升效率**。

> 有 binlog 组提交，那有 redo log 组提交吗？

这个要看 MySQL 版本，MySQL 5.6 没有 redo log 组提交，MySQL 5.7 有 redo log 组提交。

在 MySQL 5.6 的组提交逻辑中，每个事务各自执行 prepare 阶段，也就是各自将 redo log 刷盘，这样就没办法对 redo log 进行组提交。

所以在 MySQL 5.7 版本中，做了个改进，在 prepare 阶段不再让事务各自执行 redo log 刷盘操作，而是推迟到组提交的 flush 阶段，也就是说 prepare 阶段融合在了 flush 阶段。

这个优化是将 redo log 的刷盘延迟到了 flush 阶段之中，sync 阶段之前。通过延迟写 redo log 的方式，为 redolog 做了一次组写入，这样 binlog 和 redo log 都进行了优化。

接下来介绍每个阶段的过程，注意下面的过程针对的是“双 1” 配置（sync_binlog 和 innodb_flush_log_at_trx_commit 都配置为 1）。

> flush 阶段

第一个事务会成为 flush 阶段的 Leader，此时后面到来的事务都是 Follower ：

![image-20240526134121394](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526134121394.png)

接着，获取队列中的事务组，由绿色事务组的 Leader 对 redo log 做一次 write + fsync，即一次将同组事务的 redolog 刷盘：

![image-20240526134137268](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526134137268.png)

完成了 prepare 阶段后，将绿色这一组事务执行过程中产生的 binlog 写入 binlog 文件（调用 write，不会调用 fsync，所以不会刷盘，binlog 缓存在操作系统的文件系统中）。

![image-20240526134150434](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526134150434.png)

##### 总结

****

具体更新一条记录`UPDATE t_user SET name = 'xiaoming' WHERE id  = 1;`l
流程如下：

1. 执行器复制执行，会调佣存储引擎的接口，通过主键索引树搜索获取id = 1这一行记录：
   - 如果id = 1这行所在的数据页本来就在buffer pool中， 就直接返回给执行器更新；
   - 如果记录不在buffer pool，将数据页从磁盘读入到buffer pool，返回记录给执行器。
2. 执行器得到聚簇索引记录后，会看一下更新前的记录和更新后的记录是否一样：
   - 如果一样的话就不进行后续更新的流程；
   - 如果不一样的话就把更新前的记录和更新后的记录都作参数传给InnoDB层，让InnoDB真正的执行更新记录的操作；
3. 开启事务，InnoDB层更新记录前，首先要记录相应的undo log，因为这是更新操作，需要把被更新的列的旧值记下来，也就是要生成一条undo log，undo log会写入到buffer pool中的Undo页面，不过在内存修改后，需要记录对应的redo log。
4. InnoDB层开始更新记录，会先更新内存(同时标记为脏页)，然后将记录写到redo log里面，这个时候更新就算完成了，为了减少磁盘I/O，不会立即将脏页写入磁盘，后续有后台线程选择好一个合适的实践将脏页写入到磁盘，而是先写入redo日志，然后再合适的时间再修改的行数据写到磁盘上。
5. 至此，一条记录更新完了。
6. 在一条更新语句执行完成后，然后开始记录该语句对应的binlog，此时记录的binlog会被保存到binlog cache，并没有刷新到磁盘上的binlog文件，在事务提交时才会统一将该事务运行过程中的所有binlog刷新到磁盘。
7. 事务提交（为了方便说明，这里不说组提交的过程，只说两阶段提交）：
   - **prepare阶段：**将redo log对应的事务状态设置为prepare，然后将redo log刷新到磁盘；
   - **commit阶段：**将binlog刷新到磁盘，接着调用引擎提交事务接口，将redo log状态设置为commit(将事务设置为commit状态后，刷入到磁盘redo log文件)；
8. 至此，一条更新语句执行完成。

#### 揭开Buffer Pool的面纱

![image-20240526163511880](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526163511880.png)

##### 为什么要有Buffer Pool？

****

虽然说MySQL的数据是存储在磁盘里的，但是也不能每次都从磁盘里面读取数据，这样性能是极差的。

要想提升查询性能，加个缓存就行了嘛。所以，当数据从磁盘读取出来后，缓存内存中，下次查询同样是的数据时，直接从内存中读取。

为此，InnoDB存储引擎设计了一个**缓冲池（Buffer Pool）**，来提高数据库的读写性能。

![image-20240526164047314](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526164047314.png)

有了缓冲池后：

- 当读取数据时，如果数据存在于Buffer Pool中，客户端就直接读取Buffer Pool中的数据，否则再去磁盘读取。
- 当修改数据时，首先是修改Buffer Pool中数据所在的页，然后将其页设置为脏页，最后由后台线程将脏页写入到磁盘。

##### Buffer Pool有多大？

Buffer Pool是在MySQL启动的时候，想操作系统申请的一片连续的内存空间，默认配置下Buffer Pool只有**128MB。**

可以通过调整**innodb_buffer_pool_sizi**参数来设置Buffer Pool的大小，一般建议设置成可用物理内存的60%~80~。

##### Buffer Pool缓存什么？

InnoDB会把存储的数据划分为若干个[页]，也页作为磁盘和内存交互的基本单位，一个也默认大小为16KB。因此，Buffer Pool同样需要按照[页]来划分。

在MySQL启动的时候，InnoDB会为Buffer Pool申请一片连续的内存空间，然后按照默认的16KB的大小划分出一个个的页，Buffer Pool中的页就叫做缓存页。此时这些缓存页都是空闲的，之后随着程序的运行，才会有磁盘上的页被缓存到Buffer Pool中。

所以，MySQL刚启动的时候，你会观察到使用的虚拟空间很大，而使用的到的物理空间却很小，这是因为只有这些虚拟内存被访问后，操作系统才会触发缺页中断，接着将虚拟地址和物理地址建立映射关系。

Buffer Pool除了缓存[索引页]和[数据页]，还包括了undo页，插入缓存、自适应哈希索引、锁信息等等。

![image-20240526170113779](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526170113779.png)

为了更好的管理这些在Buffer Pool中的缓存页，InnoDB为每一个缓存页都创建了一个**控制块**，控制块信息包括[缓存页的表空间、页好、缓存页地址、链表节点]等等。

控制块也是占有内存空间的，它是放在Buffer Pool的最前面，接着才是缓存页，如下图：

![image-20240526170328431](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526170328431.png)

上图中控制块和缓存页之间灰色部分称为碎片空间。

> 为什么会有碎片空间呢？

你想想啊，每一个控制块都对应一个缓存页，那在分配足够多的控制块和缓存也后，可能剩余的那点空间不够一对控制块和缓存页的大小，自然就用不到咯，这个用不到的那点儿内存空间就成为碎片了。

当然，如果你把Buffer Pool的大小设置的刚刚好的话，也可能不会产生碎片。

> 查询一条记录，就需要缓存一条记录吗？

不是的。

当我们查询一条记录时，InnoDB是会把整个页的数据加载到Buffer Pool中，因为，通过索引只能定位到磁盘中的页，而不能定位到页中的一条记录。将页加载到Buffer Pool后，在通过页里的页目录去定位到某条具体的记录。

##### 如何管理空闲页？

Buffer Pool是一片连续的内存空间，当MySQL运行一段时间后，这片连续的内存空间中的缓存页既有空闲的，也有被使用的。

那当我们从磁盘读取数据的时候，总不能通过遍历这一片连续的内存空间来找到空闲的缓存页把，这样的效率太低了。所以，为了能够快速的找到空闲缓存页，可以使用链表结构，将空闲缓存页的[控制块]作为链表的节点，这个链表成为**Free链表（空闲链表）**

![image-20240526171639982](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526171639982.png)

Free链表上除了有控制块，还有一个头节点，该头节点包含链表的头节点地址，尾节点地址，以及当前链表中节点的数量等信息。

Free链表节点是一个一个的控制块，而每个控制块包含着对应缓存页的地址，所以相当于Free链表节点都对应一个空闲的缓存页。

有了Free链表后，每当需要从磁盘中加载一个页到Buffer Pool中时，就从Free链表中取一个空闲的缓存页，并且把该缓存页对应的控制块的信息填上，然后把该缓存页对应的控制块从Free链表移除。

##### 如何提高缓存命中率？

Buffer Pool的大小是有限的，对于一些频繁访问的数据我们希望可以一直留在Buffer Pool中，而一些很少访问的数据希望可以在某些时机可以淘汰掉，从而保证Buffer Pool不会因为满了而导致无法再缓存新的数据，同时还能拿保证常用数据留在Buffer Pool中。

要实现这个，最容易想到的就是LRU（Least recently used）算法。

该算法的思路是，链表头部的节点是最近使用的，而链表末尾的节点是最久没被使用的。那么，空间不够了，就淘汰最久没被使用的节点，从而腾出空间。

简单的LRU算法的实现思路是这样的：

- 当访问的页在Buffer Pool里，就直接把该页对应的LRU链表节点移动到链表的头部。
- 当访问的页不在Buffer Pool里，除了要把页放入到LRU链表的头部，还要淘汰LRU链表末尾的节点。

比如下图，假设LRU链表长度为5，LRU链表从左到右有1,2,3,4,5的页。

![image-20240526173811152](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526173811152.png)

如果访问了3号的页，因为3号页在Buffer Pool里，所以把3号页移动到头部即可。

![image-20240526173908251](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526173908251.png)

而如果接下来，访问了8号页，因为8号页不在Buffer Pool里，所以需要先淘汰末尾的5号页，然后再讲8号页加入到头部。

![image-20240526174030441](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526174030441.png)

到这里我们可以知道，Buffer Pool里有三种页和链表来管理数据。

![image-20240526174100445](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526174100445.png)

图中：

- Free Page(空闲页)，表示此页未被使用，位于Free链表；
- Clean Page（干净页），表示此页已经被使用，但是页面未发生修改，位于LRU链表。
- Dirty Page（脏页），表示此页[已被使用]且[已经被修改]，其数据和磁盘的数据已经不一致。当脏页的数据写入磁盘后，内存数据和磁盘数据，那么该页就变成了干净页。脏页同时存在于LRU链表和Flush链表。

简单的LRU算法并没有被MySQL使用，因为简单的LRU算法无法避免下面两个问题：

- 预读失效；
- Buffer Pool污染；

> 什么是预读失效？

先来说说MySQL的预读机制。程序是有空间局限性的，靠近当前被访问数据的数据，在未来大概率会被访问到。

所以，MySQL在加载数据页时，会提前把它相邻的数据页一并加载进来，目的是为了减少磁盘IO。

但是可能这些**被提前加载进来的数据页，并没有被访问，**相当于这个预读是白做了，这个就是**预读失效。**

如果使用简单是LRU算法，就会把预读页放到LRU链表头部，而当Buffer Pool空间不够的时候，还需要把末尾的页淘汰掉。

如果这些预读页一直不回被访问到，就会出现一个很奇怪的问题，不会被访问的预读页却占用了LRU链表前排的位置，而末尾淘汰的页，可能是频繁访问的页，这样就打打降低了缓存命中率。

> 怎么解决预读失效而导致缓存命中率降低的问题？

我们不能因为害怕预读失效，而将预读机制去掉，大部分情况下，局部性原理还是成立的。

要避免预读失效带来的影响，最好就是让预读的页停留在Buffer Pool里的时间要尽可能的短，让真正被访问的页才移动到LRU链表的头部，从而保证真正被读取的热数据留在Buffer Pool里的时间尽可能长。

那到底怎么才能避免呢？

MySQL是这样做的，它改进了LRU算法，将LRU划分为2个区域:**old区域**和**young区域**。

young区域在LRU链表的前半部分，old区域则是在后半部分，如下图：

![image-20240526202758345](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526202758345.png)

old区域占整个LRU链表长度比例可以通过**innodb_old_blocks_pct**参数来设置，默认是37，代表整个LRU链表中young区域与ol区域比例是63:37。

**划分这两个区域后，预读的页就只需要加入到old区域的头部，当页被真正访问的时候，才将页插入young区域头部。如果预读的页一直没有被真正访问，就会从old区域移除，这样就不会影响young区域中的热点数据。**

接下来，下面举个例子：

假设有一个长度为10的LRU的链表，其中young区域占比70%，old区域占比30%。

![image-20240526203413543](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526203413543.png)

现在有一个编号为20的页被预读了，这个页只会被插入到old区域头部，而old区域末尾的页（10号）会被淘汰掉。

![image-20240526203546788](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526203546788.png)

如果20号页一直不回被访问，它也没有占用到young区域的位置，而且还会比young区域的数据更早被淘汰出去。

如果20号页被预读后，立刻被访问了，那么就会将它插入到了young区域的头部，young区域末尾的页（7号），会被挤到old区域，作为old区域的头部，这个过程并不会有页被淘汰。

![image-20240526203809807](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526203809807.png)

虽然通过划分old区域和young区域避免了预读失效带来的影响，但是还有个问题无法解决，那就是Buffer Pool污染的问题。

> 什么是Buffer Pool污染？

当某一个SQL语句扫描了大量的数据时，**在Buffer Pool空间比价有限的情况下，可能会将Buffer Pool里的所有页都替换出去，导致大量热数据被淘汰了**，等这些热数据又被再次访问的时候，由于缓存未命中，就会产生大量的磁盘I/O，MySQL性能就会急剧下降，这个过程称为**Buffer Pool污染。**

注意，Buffer Pool污染并不只是查询语句查询出了大量的数据才出现的问题，即使查询出来的结果集很小，也会造成Buffer Pool污染。

比如，在一个数据量非常大的表，执行了这条语句：

```
select * from t_user where name like "%xiaoming%";
```

可能这个查询的结构就几条记录，但是由于这条语句会发生索引失效，所以这个查询过程是全表扫描的，接着发生如下的过程：

- 从磁盘读到的页加到LRU链表的old区域头部；
- 当从页里读取行记录时，也就是页被访问的时候，就要将该页放发哦young区域头部；
- 接下来拿行记录的name字段和字符串xiaoming进行模糊匹配，如果符合条件，就加到结果集里；
- 如此往复，直到扫描完表中的所有记录。

经过这一番折腾，原本young区域的热点数据都会被替换掉。

举个例子，假设需要批量扫描：21,22,23,24,25这五个页，这些页都会被逐一访问（读取页里的记录）。

![image-20240526205823631](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526205823631.png)

在批量访问这些数据的时候，会被逐一插入到 young 区域头部。

![image-20240526210100887](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240526210100887.png)

可以看到，原本在 young 区域的热点数据 6 和 7 号页都被淘汰了，这就是 Buffer Pool 污染的问题。

> 怎么解决出现Buffer Pool污染而导致缓存命中率下降问题？

像前面这种全表扫描的查询，很多缓冲页其实只会被访问一次，但是它却只因为被访问一次而进入到young区域，从而导致热点数据被替换了。

LRU链表中young区域就是热点数据，只要我们提高进入到young的门槛，就能有效地保证young区域里的热点数据不会被替换掉。MySQL是这样做的，进入到young区域条件增加了了一个**停留在old区域的时间判断**。

具体是这样做的，在对某个处在old区域的缓存页进行第一场访问时，就在它对应控制块中记录下来这个访问时间：

- 如果后续的访问与第一次访问的时间**在某个时间间隔内**，那么该缓存页就不会被从old区域移动到young区域的头部；
- 如果后续的访问时间与第一访问的时间**不在某个时间间隔内**，那么该缓存页移动到young区域的头部；

也就是说，**只有同时满足[被访问]与[在old区域停留时间超过1秒]两个条件，才会被插入到young区域头部，**这样就解决了Buffer Pool污染的问题。

另外，MySQL针对young区域其实做了一个优化，为了防止young区域节点频繁移动到头部。young区域前1/4被访问不会移动到链表头部，只有后面的3/4被访问了才会。

##### 脏页什么时候会被刷入磁盘？

引入了Buffer Pool后，当修改数据时，首先是修改了Buffer Pool中数据所在的页，然后将其页设置为脏页，但是磁盘中还是原数据。因此，脏页需要被刷入磁盘，保证缓存和磁盘数据一致，但是若每次修改都要刷入磁盘，则性能会很差，因此一般都会在一定时机进行批量刷盘。

可能大家担心，如果在脏页还没来得及刷入到磁盘时，MySQL宕机了，不就丢失数据了嘛？

这个不用担心，InnoDB的更新操作采用的是Write Ahead Log策略，即先写日志，再写磁盘，通过redo Log让日志MySQL拥有了崩溃恢复能力。

下面几种情况会出发脏页的刷新：

- 当redo log日志满了的情况下，会主动触发脏页刷新的磁盘；
- Buffer Pool空间不足时，需要将一部分数据页淘汰掉，如果淘汰的是脏页，需要先将脏页同步到磁盘；
- MySQL正常关闭之前，会把所有的脏页刷入到磁盘；

在我们开启了慢SQL监控后，如果你发现[偶尔]会出现一些用时稍长的SQL，这可能是因为脏页再刷新到磁盘时可能会给数据库带来性能开销，导致数据库操作抖动。

如果间断出现这种现象，就需要调大Buffer Pool空间或redo log日志的大小。

##### 总结

****

InnoDB存储引擎设计了一个**缓冲池（Buffer Pool）**，来提高数据库的读写性能。

Buffer Pool以页为单位缓冲数据，可以通过**innodb_buffer_pool_size**参数调整缓冲池的大小，默认是128MB。

InnoDB通过三种链表来管理缓存页：

- Free List（空闲页链表），管理空闲页；
- Flush List（脏页链表），管理脏页；
- LRU List，管理脏页+干净页，将最近且经常查询的数据缓存在其中，而不经常查询的数据就淘汰出去；

InnoDB 对 LRU 做了一些优化，我们熟悉的 LRU 算法通常是将最近查询的数据放到 LRU 链表的头部，而 InnoDB 做 2 点优化：

- 将 LRU 链表 分为**young 和 old 两个区域**，加入缓冲池的页，优先插入 old 区域；页被访问时，才进入 young 区域，目的是为了解决预读失效的问题。
- 当**「页被访问」且「 old 区域停留时间超过 `innodb_old_blocks_time` 阈值（默认为1秒）」**时，才会将页插入到 young 区域，否则还是插入到 old 区域，目的是为了解决批量数据访问，大量热数据淘汰的问题。

可以通过调整 `innodb_old_blocks_pct` 参数，设置 young 区域和 old 区域比例。

在开启了慢 SQL 监控后，如果你发现「偶尔」会出现一些用时稍长的 SQL，这可因为脏页在刷新到磁盘时导致数据库性能抖动。如果在很短的时间出现这种现象，就需要调大 Buffer Pool 空间或 redo log 日志的大小。





