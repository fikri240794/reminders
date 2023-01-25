package persistences

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func findAllReminderQuery(driver string, dto *dtos.ReminderPaginationParamsDTO) (string, []interface{}, error) {
	var (
		query string
		args  []interface{}
		err   error
		exps  []exp.Expression
	)

	if dto == nil {
		dto = dtos.NewReminderPaginationParamsDTO()
	}

	exps = []exp.Expression{goqu.C("is_deleted").IsFalse()}

	if dto.Keyword != "" {
		exps = append(exps, goqu.C("note").ILike(fmt.Sprintf("%%%s%%", dto.Keyword)))
	}

	query, args, err = goqu.Dialect(driver).
		From("reminders").
		Select(
			"id",
			"note",
			"created_at",
			"updated_at",
			"is_deleted",
		).
		Where(goqu.And(exps...)).
		Order(
			goqu.C("created_at").Asc(),
			goqu.C("updated_at").Asc(),
		).
		Limit(uint(dto.Take)).
		Offset(uint(dto.Skip)).
		Prepared(true).
		ToSQL()

	if err != nil {
		return "", nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	return query, args, nil
}

func countAllReminderQuery(driver string, dto *dtos.ReminderPaginationParamsDTO) (string, []interface{}, error) {
	var (
		query string
		args  []interface{}
		err   error
		exps  []exp.Expression
	)

	exps = []exp.Expression{goqu.C("is_deleted").IsFalse()}

	if dto.Keyword != "" {
		exps = append(exps, goqu.C("note").ILike(fmt.Sprintf("%%%s%%", dto.Keyword)))
	}

	query, args, err = goqu.Dialect(driver).
		From("reminders").
		Select(
			goqu.COUNT("id").
				As("count"),
		).
		Where(goqu.And(exps...)).
		Prepared(true).
		ToSQL()

	if err != nil {
		return "", nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	return query, args, nil
}

func findOneReminderQuery(driver string, keyword string) (string, []interface{}, error) {
	var (
		query string
		args  []interface{}
		err   error
		exps  []exp.Expression
	)

	exps = []exp.Expression{goqu.C("is_deleted").IsFalse()}

	if keyword != "" {
		exps = append(exps, goqu.C("note").ILike(fmt.Sprintf("%%%s%%", keyword)))
	}

	query, args, err = goqu.Dialect(driver).
		From("reminders").
		Select(
			"id",
			"note",
			"created_at",
			"updated_at",
			"is_deleted",
		).
		Where(goqu.And(exps...)).
		Order(
			goqu.C("created_at").Asc(),
			goqu.C("updated_at").Asc(),
		).
		Limit(1).
		Offset(0).
		Prepared(true).
		ToSQL()

	if err != nil {
		return "", nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	return query, args, nil
}

func findReminderByIdQuery(driver string, id string) (string, []interface{}, error) {
	var (
		query string
		args  []interface{}
		err   error
	)

	query, args, err = goqu.Dialect(driver).
		From("reminders").
		Select(
			"id",
			"note",
			"created_at",
			"updated_at",
			"is_deleted",
		).
		Where(
			goqu.C("id").Eq(id),
			goqu.C("is_deleted").IsFalse(),
		).
		Prepared(true).
		ToSQL()

	if err != nil {
		return "", nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	return query, args, nil
}

func createReminderQuery(driver string, e *entities.ReminderEntity) (string, []interface{}, error) {
	var (
		query string
		args  []interface{}
		err   error
	)

	if e == nil {
		return "", nil, errors.NewError(base_enums.BAD_REQUEST, "reminder entity is null")
	}

	query, args, err = goqu.Dialect(driver).
		Insert("reminders").
		Rows(
			goqu.Record{
				"id":         e.ID,
				"note":       e.Note,
				"created_at": e.CreatedAt,
				"updated_at": e.UpdatedAt,
				"is_deleted": e.IsDeleted,
			},
		).
		Prepared(true).
		ToSQL()

	if err != nil {
		return "", nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	return query, args, nil
}

func updateReminderQuery(driver string, e *entities.ReminderEntity) (string, []interface{}, error) {
	var (
		query string
		args  []interface{}
		err   error
	)

	if e == nil {
		return "", nil, errors.NewError(base_enums.BAD_REQUEST, "reminder entity is null")
	}

	query, args, err = goqu.Dialect(driver).
		Update("reminders").
		Set(
			goqu.Record{
				"note":       e.Note,
				"updated_at": e.UpdatedAt,
				"is_deleted": e.IsDeleted,
			},
		).
		Where(
			goqu.C("id").Eq(e.ID),
		).
		Prepared(true).
		ToSQL()

	if err != nil {
		return "", nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	return query, args, nil
}
