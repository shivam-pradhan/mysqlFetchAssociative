# mysqlFetchAssociative

Project to get array of  map where column_name is key and its attribute_data is its value



=> => => => => => => => => => => => => => => => => => => => => => => => => => => => 

Steps to run project:

1) clone the repo

2) change your DB name, table, username & password in main.go

3) write any select query for associative fetch

4) run main.go



=> => => => => => => => => => => => => => => => => => => => => => => => => => => => 





Below is a bsic example of project output on certain DB.

/******* TEST TABLE STRUCTURE **************

	mysql> desc testTable;

+----------------+--------------+------+-----+---------+-------+

| Field          | Type         | Null | Key | Default | Extra |

+----------------+--------------+------+-----+---------+-------+

| id             | int(11)      | YES  |     | NULL    |       |

| full_names     | varchar(150) | NO   |     | NULL    |       |

| gender         | varchar(6)   | YES  |     | NULL    |       |

| date_of_birth  | date         | YES  |     | NULL    |       |

| contact_number | varchar(75)  | YES  |     | NULL    |       |

| email          | varchar(255) | YES  |     | NULL    |       |

+----------------+--------------+------+-----+---------+-------+

	*************** TABLE VALUES ******************

	mysql> select * from testTable;

+------+------------+--------+---------------+----------------+------------------+

| id   | full_names | gender | date_of_birth | contact_number | email            |

+------+------------+--------+---------------+----------------+------------------+

|    1 | testname1  | male   | 2019-01-01    | 9909123456     | example1@abc.com |

|    2 | testname2  | male   | NULL          | 9909123457     | example2@abc.com |

|    3 | testname3  | male   | 2019-01-02    | 9909123456     | NULL             |

+------+------------+--------+---------------+----------------+------------------+

	*************** ARRAY OF MAP **********************

[

	map[

		contact_number:9909123456

		date_of_birth:2019-01-01 00:00:00 +0000 UTC

		email:example1@abc.com

		full_names:testname1

		gender:male

		id:1

	]

	map[

		contact_number:9909123457

		date_of_birth:<nil>

		email:example2@abc.com

		full_names:testname2

		gender:male

		id:2

	]

	map[

		contact_number:9909123456

		date_of_birth:2019-01-02 00:00:00 +0000 UTC

		email:

		full_names:testname3

		gender:male

		id:3

	]

]

	***********************/

