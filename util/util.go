/*
 * @Author: small_ant xms.chnb@gmail.com
 * @Time: 2022-11-17 14:06:27
 * @LastAuthor: small_ant xms.chnb@gmail.com
 * @lastTime: 2022-11-18 09:46:02
 * @FileName: util
 * @Desc:
 *
 * Copyright (c) 2022 by small_ant xms.chnb@gmail.com, All Rights Reserved.
 */

package util

import (
    "math/big"
    "strings"

    "github.com/shopspring/decimal"
)

func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
    value := new(big.Int)
    switch v := ivalue.(type) {
    case string:
        value.SetString(v, 10)
    case *big.Int:
        value = v
    }

    mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
    num, _ := decimal.NewFromString(value.String())
    result := num.Div(mul)

    return result
}

// Convert the two strings to decimal numbers, add them, and convert the result back to a string.
//
// Args:
//   a (string): The first number to add.
//   b (string): The base to use for representing a numeric value. Must be between 2 and 36 inclusive,
// or 0.
//
// Returns:
//   The sum of the two decimal numbers.
func SumDecimal(a, b string) string {
    c, _ := decimal.NewFromString(a)
    d, _ := decimal.NewFromString(b)
    return c.Add(d).String()
}

// Subtract two decimal numbers represented as strings.
//
// Args:
//   a (string): The first number to be subtracted
//   b (string): the base number
//
// Returns:
//   The difference between the two numbers.
func SubDecimal(a, b string) string {
    c, _ := decimal.NewFromString(a)
    d, _ := decimal.NewFromString(b)
    return c.Sub(d).String()
}

// It divides two decimal numbers.
//
// Args:
//   a (string): dividend
//   b (string): the base number
//
// Returns:
//   The result of the division of the two decimal numbers.
func DivDecimal(a, b string) string {
    c, _ := decimal.NewFromString(a)
    d, _ := decimal.NewFromString(b)
    return DecimalInt(c.Div(d))
}

// It takes a decimal.Decimal and returns a string of the integer part of the decimal
//
// Args:
//   d: The decimal.Decimal object you want to convert to an integer.
//
// Returns:
//   The integer part of the decimal.
func DecimalInt(d decimal.Decimal) string {
    return strings.Split(d.String(), ".")[0]
}

// Compare returns -1 if a < b, 0 if a == b, and 1 if a > b.
//
// Args:
//   a (string): The first number to compare.
//   b (string): the base to use for representing a numeric value. Must be between 2 and 36 inclusive,
// or 0.
//
// Returns:
//   The return value is the result of the comparison between the two decimal values.
func Compare(a, b string) int {
    c, _ := decimal.NewFromString(a)
    d, _ := decimal.NewFromString(b)
    return c.Cmp(d)
}
