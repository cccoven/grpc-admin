package util

import (
	"errors"
	"github.com/jinzhu/copier"
	"time"
)

func CopyWithTimeConverter(dst any, src any) error {
	return copier.CopyWithOption(dst, src, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},       // 例如: model 的 time.Time 类型，如 CreatedAt
				DstType: int64(copier.Int), // 例如：proto 文件的 int64（时间戳）类型
				Fn: func(src any) (any, error) {
					s, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return s.UnixMilli(), nil
				},
			},
			{
				SrcType: int64(copier.Int), // 例如：proto 文件的 int64（时间戳）类型
				DstType: copier.String,     // 例如：返回到前端的时间格式字符串
				Fn: func(src any) (any, error) {
					s, ok := src.(int64)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					t := time.UnixMilli(s)
					return t.Format("2006-01-02 15:04:05"), nil
				},
			},
		},
	})
}
