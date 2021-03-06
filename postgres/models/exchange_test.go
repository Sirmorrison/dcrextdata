// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testExchanges(t *testing.T) {
	t.Parallel()

	query := Exchanges()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testExchangesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExchangesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Exchanges().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExchangesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ExchangeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExchangesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ExchangeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Exchange exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ExchangeExists to return true, but got false.")
	}
}

func testExchangesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	exchangeFound, err := FindExchange(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if exchangeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testExchangesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Exchanges().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testExchangesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Exchanges().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testExchangesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	exchangeOne := &Exchange{}
	exchangeTwo := &Exchange{}
	if err = randomize.Struct(seed, exchangeOne, exchangeDBTypes, false, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}
	if err = randomize.Struct(seed, exchangeTwo, exchangeDBTypes, false, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = exchangeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = exchangeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Exchanges().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testExchangesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	exchangeOne := &Exchange{}
	exchangeTwo := &Exchange{}
	if err = randomize.Struct(seed, exchangeOne, exchangeDBTypes, false, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}
	if err = randomize.Struct(seed, exchangeTwo, exchangeDBTypes, false, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = exchangeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = exchangeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testExchangesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExchangesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(exchangeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExchangeToManyExchangeTicks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Exchange
	var b, c ExchangeTick

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, exchangeTickDBTypes, false, exchangeTickColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, exchangeTickDBTypes, false, exchangeTickColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ExchangeID = a.ID
	c.ExchangeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ExchangeTicks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ExchangeID == b.ExchangeID {
			bFound = true
		}
		if v.ExchangeID == c.ExchangeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ExchangeSlice{&a}
	if err = a.L.LoadExchangeTicks(ctx, tx, false, (*[]*Exchange)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ExchangeTicks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ExchangeTicks = nil
	if err = a.L.LoadExchangeTicks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ExchangeTicks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testExchangeToManyAddOpExchangeTicks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Exchange
	var b, c, d, e ExchangeTick

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ExchangeTick{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, exchangeTickDBTypes, false, strmangle.SetComplement(exchangeTickPrimaryKeyColumns, exchangeTickColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ExchangeTick{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddExchangeTicks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ExchangeID {
			t.Error("foreign key was wrong value", a.ID, first.ExchangeID)
		}
		if a.ID != second.ExchangeID {
			t.Error("foreign key was wrong value", a.ID, second.ExchangeID)
		}

		if first.R.Exchange != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Exchange != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ExchangeTicks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ExchangeTicks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ExchangeTicks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testExchangesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testExchangesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ExchangeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testExchangesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Exchanges().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	exchangeDBTypes = map[string]string{`ID`: `integer`, `Name`: `text`, `URL`: `text`}
	_               = bytes.MinRead
)

func testExchangesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(exchangePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(exchangeAllColumns) == len(exchangePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testExchangesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(exchangeAllColumns) == len(exchangePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Exchange{}
	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, exchangeDBTypes, true, exchangePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(exchangeAllColumns, exchangePrimaryKeyColumns) {
		fields = exchangeAllColumns
	} else {
		fields = strmangle.SetComplement(
			exchangeAllColumns,
			exchangePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ExchangeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testExchangesUpsert(t *testing.T) {
	t.Parallel()

	if len(exchangeAllColumns) == len(exchangePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Exchange{}
	if err = randomize.Struct(seed, &o, exchangeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Exchange: %s", err)
	}

	count, err := Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, exchangeDBTypes, false, exchangePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Exchange: %s", err)
	}

	count, err = Exchanges().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
