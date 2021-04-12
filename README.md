# DT

这是一套神奇的数据结构，用于在 Go 范型出现之前统一各种数据类型的表达

本包中定义了若干组 interface，并遵循如下规范

+ interface 名称为 Interface，则其一定实现了 AsInterface 方法
+ 子 Interface 一定实现了父 Interface 的方法

+ Number 所有数字
    + Integer 所有整数
      + Int 有符号整数
        + Int8
        + Int16
        + Int32
        + Int64
        + Int128
      + UInt 无符号整数
        + UInt8
        + UInt16
        + UInt32
        + UInt64
        + UInt128
    + Float 所有浮点数
    + BigDecimal (TODO)