package mysqlOrm

import "github.com/naufalfmm/csv-transaction/utils/orm"

type SelectQueryModifier = func(so orm.Select) orm.Select
type UpdateQueryModifier = func(uq orm.Update) orm.Update
