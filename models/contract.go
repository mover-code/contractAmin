package models

import (
    "gorm.io/plugin/dbresolver"
)

type Contract struct {
    Id       int64
    Addr     string
    ChainRpc string `gorm:"column:chainRpc"`
    Abi      string
    Scan     string
    ChainId  string `gorm:"column:chainId"`
}

func FirstContract() *Contract {
    s := new(Contract)
    // orm.Clauses(dbresolver.Write).First(s)
    // orm.Clauses(dbresolver.Use("blind")).First(s)
    return s
}

func (s *Contract) ContractInfo(addr string) *Contract {
    orm.Clauses(dbresolver.Write).Model(Contract{}).Where("addr like ?", addr).First(&s)
    return s
}
