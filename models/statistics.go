package models

import (
	"fmt"
	"html/template"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Statistics struct {
	ID         uint `gorm:"primary_key,column:cpu"`
	CPU        uint `gorm:"column:cpu"`
	Likes      uint `gorm:"column:likes"`
	Sales      uint `gorm:"column:sales"`
	NewMembers uint `gorm:"column:new_members"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type BlindBoxStatistics struct {
	DB *gorm.DB
}

func FirstStatics() *Statistics {
	s := new(Statistics)
	orm.Clauses(dbresolver.Write).First(s)
	// orm.Clauses(dbresolver.Use("blind")).First(s)
	return s
}

func (s *Statistics) CPUTmpl() template.HTML {
	return template.HTML(strconv.Itoa(int(s.CPU)))
}

func (s *Statistics) LikesTmpl() template.HTML {
	return template.HTML(strconv.Itoa(int(s.Likes)))
}

func (s *Statistics) SalesTmpl() template.HTML {
	return template.HTML(strconv.Itoa(int(s.Sales)))
}

func (s *Statistics) NewMembersTmpl() template.HTML {
	return template.HTML(strconv.Itoa(int(s.NewMembers)))
}

// 统计盲盒售卖情况

func Analysis() *BlindBoxStatistics {
	return &BlindBoxStatistics{
		DB: orm.Clauses(dbresolver.Use("blind")),
	}
}

// A function that is used to count the number of sales of the blind box.
func (b *BlindBoxStatistics) SumSalesTmpl() template.HTML {
	sum := 0
	// b.DB.Raw("select count(*) from box_order").Scan(&sum)
	return template.HTML(strconv.Itoa(int(sum)))
}

// Counting the amount of money that the blind box has earned.
func (b *BlindBoxStatistics) AmountSalesTmpl() template.HTML {
	var sum float64
	// b.DB.Raw("select sum(buyer_value) from box_order").Scan(&sum)
	return template.HTML(fmt.Sprintf("%.2f BNB", sum))
}

// Counting the number of today sales of the blind box.
func (b *BlindBoxStatistics) ToDaySumSalesTmpl() template.HTML {
	// t := time.Now()
	// addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	sum := 0
	// b.DB.Raw(fmt.Sprintf("select count(*) from box_order where create_time >= %v", addTime.Unix())).Scan(&sum)
	return template.HTML(strconv.Itoa(int(sum)))
}

// Counting the number of users who have purchased the blind box.
func (b *BlindBoxStatistics) UserSumTmpl() template.HTML {
	sum := 0
	// b.DB.Raw("select count(distinct owner) from box_order").Scan(&sum)
	return template.HTML(strconv.Itoa(int(sum)))
}

// A function that is used to count the amount of money that the blind box has earned today.
func (b *BlindBoxStatistics) TodayAmountSalesTmpl() template.HTML {
	// t := time.Now()
	// addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	var sum float64
	// Getting the current time.
	// b.DB.Raw("select sum(buyer_value) from box_order where create_time >= ?", addTime.Unix()).Scan(&sum)
	return template.HTML(fmt.Sprintf("%v BNB", sum))
}
