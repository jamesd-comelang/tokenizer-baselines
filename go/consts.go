package main

const (
	TK_SEMI = 1
	TK_EXPLAIN = 2
	TK_QUERY = 3
	TK_PLAN = 4
	TK_BEGIN = 5
	TK_TRANSACTION = 6
	TK_DEFERRED = 7
	TK_IMMEDIATE = 8
	TK_EXCLUSIVE = 9
	TK_COMMIT = 10
	TK_END = 11
	TK_ROLLBACK = 12
	TK_SAVEPOINT = 13
	TK_RELEASE = 14
	TK_TO = 15
	TK_TABLE = 16
	TK_CREATE = 17
	TK_IF = 18
	TK_NOT = 19
	TK_EXISTS = 20
	TK_TEMP = 21
	TK_LP = 22
	TK_RP = 23
	TK_AS = 24
	TK_COMMA = 25
	TK_WITHOUT = 26
	TK_ABORT = 27
	TK_ACTION = 28
	TK_AFTER = 29
	TK_ANALYZE = 30
	TK_ASC = 31
	TK_ATTACH = 32
	TK_BEFORE = 33
	TK_BY = 34
	TK_CASCADE = 35
	TK_CAST = 36
	TK_CONFLICT = 37
	TK_DATABASE = 38
	TK_DESC = 39
	TK_DETACH = 40
	TK_EACH = 41
	TK_FAIL = 42
	TK_OR = 43
	TK_AND = 44
	TK_IS = 45
	TK_ISNOT = 46
	TK_MATCH = 47
	TK_LIKE_KW = 48
	TK_BETWEEN = 49
	TK_IN = 50
	TK_ISNULL = 51
	TK_NOTNULL = 52
	TK_NE = 53
	TK_EQ = 54
	TK_GT = 55
	TK_LE = 56
	TK_LT = 57
	TK_GE = 58
	TK_ESCAPE = 59
	TK_ID = 60
	TK_COLUMNKW = 61
	TK_DO = 62
	TK_FOR = 63
	TK_IGNORE = 64
	TK_INITIALLY = 65
	TK_INSTEAD = 66
	TK_NO = 67
	TK_KEY = 68
	TK_OF = 69
	TK_OFFSET = 70
	TK_PRAGMA = 71
	TK_RAISE = 72
	TK_RECURSIVE = 73
	TK_REPLACE = 74
	TK_RESTRICT = 75
	TK_ROW = 76
	TK_ROWS = 77
	TK_TRIGGER = 78
	TK_VACUUM = 79
	TK_VIEW = 80
	TK_VIRTUAL = 81
	TK_WITH = 82
	TK_NULLS = 83
	TK_FIRST = 84
	TK_LAST = 85
	TK_CURRENT = 86
	TK_FOLLOWING = 87
	TK_PARTITION = 88
	TK_PRECEDING = 89
	TK_RANGE = 90
	TK_UNBOUNDED = 91
	TK_EXCLUDE = 92
	TK_GROUPS = 93
	TK_OTHERS = 94
	TK_TIES = 95
	TK_GENERATED = 96
	TK_ALWAYS = 97
	TK_MATERIALIZED = 98
	TK_REINDEX = 99
	TK_RENAME = 100
	TK_CTIME_KW = 101
	TK_ANY = 102
	TK_BITAND = 103
	TK_BITOR = 104
	TK_LSHIFT = 105
	TK_RSHIFT = 106
	TK_PLUS = 107
	TK_MINUS = 108
	TK_STAR = 109
	TK_SLASH = 110
	TK_REM = 111
	TK_CONCAT = 112
	TK_PTR = 113
	TK_COLLATE = 114
	TK_BITNOT = 115
	TK_ON = 116
	TK_INDEXED = 117
	TK_STRING = 118
	TK_JOIN_KW = 119
	TK_CONSTRAINT = 120
	TK_DEFAULT = 121
	TK_NULL = 122
	TK_PRIMARY = 123
	TK_UNIQUE = 124
	TK_CHECK = 125
	TK_REFERENCES = 126
	TK_AUTOINCR = 127
	TK_INSERT = 128
	TK_DELETE = 129
	TK_UPDATE = 130
	TK_SET = 131
	TK_DEFERRABLE = 132
	TK_FOREIGN = 133
	TK_DROP = 134
	TK_UNION = 135
	TK_ALL = 136
	TK_EXCEPT = 137
	TK_INTERSECT = 138
	TK_SELECT = 139
	TK_VALUES = 140
	TK_DISTINCT = 141
	TK_DOT = 142
	TK_FROM = 143
	TK_JOIN = 144
	TK_USING = 145
	TK_ORDER = 146
	TK_GROUP = 147
	TK_HAVING = 148
	TK_LIMIT = 149
	TK_WHERE = 150
	TK_RETURNING = 151
	TK_INTO = 152
	TK_NOTHING = 153
	TK_FLOAT = 154
	TK_BLOB = 155
	TK_INTEGER = 156
	TK_VARIABLE = 157
	TK_CASE = 158
	TK_WHEN = 159
	TK_THEN = 160
	TK_ELSE = 161
	TK_INDEX = 162
	TK_ALTER = 163
	TK_ADD = 164
	TK_WINDOW = 165
	TK_OVER = 166
	TK_FILTER = 167
	TK_COLUMN = 168
	TK_AGG_FUNCTION = 169
	TK_AGG_COLUMN = 170
	TK_TRUEFALSE = 171
	TK_FUNCTION = 172
	TK_UPLUS = 173
	TK_UMINUS = 174
	TK_TRUTH = 175
	TK_REGISTER = 176
	TK_VECTOR = 177
	TK_SELECT_COLUMN = 178
	TK_IF_NULL_ROW = 179
	TK_ASTERISK = 180
	TK_SPAN = 181
	TK_ERROR = 182
	TK_QNUMBER = 183
	TK_SPACE = 184
	TK_COMMENT = 185
	TK_ILLEGAL = 186
	CC_X = 0
	CC_KYWD0 = 1
	CC_KYWD = 2
	CC_DIGIT = 3
	CC_DOLLAR = 4
	CC_VARALPHA = 5
	CC_VARNUM = 6
	CC_SPACE = 7
	CC_QUOTE = 8
	CC_QUOTE2 = 9
	CC_PIPE = 10
	CC_MINUS = 11
	CC_LT = 12
	CC_GT = 13
	CC_EQ = 14
	CC_BANG = 15
	CC_SLASH = 16
	CC_LP = 17
	CC_RP = 18
	CC_SEMI = 19
	CC_PLUS = 20
	CC_STAR = 21
	CC_PERCENT = 22
	CC_COMMA = 23
	CC_AND = 24
	CC_TILDA = 25
	CC_DOT = 26
	CC_ID = 27
	CC_ILLEGAL = 28
	CC_NUL = 29
	CC_BOM = 30
)

var kwMap = map[string]int{
	"abort": 27,
	"action": 28,
	"add": 164,
	"after": 29,
	"all": 136,
	"alter": 163,
	"always": 97,
	"analyze": 30,
	"and": 44,
	"as": 24,
	"asc": 31,
	"attach": 32,
	"autoincrement": 127,
	"before": 33,
	"begin": 5,
	"between": 49,
	"by": 34,
	"cascade": 35,
	"case": 158,
	"cast": 36,
	"check": 125,
	"collate": 114,
	"column": 61,
	"commit": 10,
	"conflict": 37,
	"constraint": 120,
	"create": 17,
	"cross": 119,
	"current": 86,
	"current_date": 101,
	"current_time": 101,
	"current_timestamp": 101,
	"database": 38,
	"default": 121,
	"deferrable": 132,
	"deferred": 7,
	"delete": 129,
	"desc": 39,
	"detach": 40,
	"distinct": 141,
	"do": 62,
	"drop": 134,
	"each": 41,
	"else": 161,
	"end": 11,
	"escape": 59,
	"except": 137,
	"exclude": 92,
	"exclusive": 9,
	"exists": 20,
	"explain": 2,
	"fail": 42,
	"filter": 167,
	"first": 84,
	"following": 87,
	"for": 63,
	"foreign": 133,
	"from": 143,
	"full": 119,
	"generated": 96,
	"glob": 48,
	"group": 147,
	"groups": 93,
	"having": 148,
	"if": 18,
	"ignore": 64,
	"immediate": 8,
	"in": 50,
	"index": 162,
	"indexed": 117,
	"initially": 65,
	"inner": 119,
	"insert": 128,
	"instead": 66,
	"intersect": 138,
	"into": 152,
	"is": 45,
	"isnull": 51,
	"join": 144,
	"key": 68,
	"last": 85,
	"left": 119,
	"like": 48,
	"limit": 149,
	"match": 47,
	"materialized": 98,
	"natural": 119,
	"no": 67,
	"not": 19,
	"nothing": 153,
	"notnull": 52,
	"null": 122,
	"nulls": 83,
	"of": 69,
	"offset": 70,
	"on": 116,
	"or": 43,
	"order": 146,
	"others": 94,
	"outer": 119,
	"over": 166,
	"partition": 88,
	"plan": 4,
	"pragma": 71,
	"preceding": 89,
	"primary": 123,
	"query": 3,
	"raise": 72,
	"range": 90,
	"recursive": 73,
	"references": 126,
	"regexp": 48,
	"reindex": 99,
	"release": 14,
	"rename": 100,
	"replace": 74,
	"restrict": 75,
	"returning": 151,
	"right": 119,
	"rollback": 12,
	"row": 76,
	"rows": 77,
	"savepoint": 13,
	"select": 139,
	"set": 131,
	"table": 16,
	"temp": 21,
	"temporary": 21,
	"then": 160,
	"ties": 95,
	"to": 15,
	"transaction": 6,
	"trigger": 78,
	"unbounded": 91,
	"union": 135,
	"unique": 124,
	"update": 130,
	"using": 145,
	"vacuum": 79,
	"values": 140,
	"view": 80,
	"virtual": 81,
	"when": 159,
	"where": 150,
	"window": 165,
	"with": 82,
	"without": 26,
}
