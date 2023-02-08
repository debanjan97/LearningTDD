package main

import "testing"

// Chapter 1: The Money Example
func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{
		amount:   5,
		currency: "USD",
	}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got %d", tenner.amount)
	}
	if tenner.currency != "USD" {
		t.Errorf("Expected USD, got %s", tenner.currency)
	}
}

type Dollar struct {
	amount int
}

func (d *Dollar) Times(multiplier int) Dollar {
	return Dollar{
		amount: d.amount * multiplier,
	}
}

// Chapter 2: Multiple Currencies
func TestMultiplicationsInEuros(t *testing.T) {
	tenEuros := Money{
		amount:   10,
		currency: "EUR",
	}
	twentyEuros := tenEuros.Times(2)
	if twentyEuros.amount != 20 {
		t.Errorf("Expected 20, got %d", twentyEuros.amount)
	}
	if twentyEuros.currency != "EUR" {
		t.Errorf("Expected EUR, got %s", twentyEuros.currency)
	}
}

type Money struct {
	amount   int
	currency string
}

func (m *Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

func TestMultiplicationsGenerics(t *testing.T) {
	tenEuros := Money{
		amount:   10,
		currency: "EUR",
	}
	twentyEuros := tenEuros.Times(2)
	if twentyEuros.amount != 20 {
		t.Errorf("Expected 20, got %d", twentyEuros.amount)
	}
	if twentyEuros.currency != "EUR" {
		t.Errorf("Expected EUR, got %s", twentyEuros.currency)
	}
}
