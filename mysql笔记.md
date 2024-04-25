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

1. **id：**select查询
2. 

- type
- key
- key_len
- rows
- filtered
- extra

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

#### 事务隔离级别是怎么实现的

##### 什么是脏读

##### 什么是幻读

##### 什么是不可重复读

##### 这几个隔离级别，分别解决了什么问题

##### MVCC是什么

- 快照读
- 当前读
- readview是什么
- rr和rc隔离级别下，readview的生成时机
- undolog版本链怎么生成的
- readview和undolog怎么比较的——四个判断依据
- mcvv能完全解决幻读问题吗

#### 可重复读隔离级别，完全解决了幻读吗

