package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/reinhardjs/mezink/domain"
	"github.com/reinhardjs/mezink/domain/dto/request"
	"github.com/reinhardjs/mezink/record/repository"
	"github.com/sirupsen/logrus"
)

type mysqlRecordRepository struct {
	Conn *sql.DB
}

func NewMysqlRecordRepository(conn *sql.DB) domain.RecordRepository {
	return &mysqlRecordRepository{conn}
}

func (m *mysqlRecordRepository) GetBySumRange(ctx context.Context, request *request.GetRecordRequest) (result []domain.Record, err error) {

	query := fmt.Sprintf(`
		SELECT *
		FROM records
		WHERE (
			SELECT SUM(value)
			FROM JSON_TABLE(
				records.marks,
				"$[*]"
				COLUMNS (
					value INT PATH "$"
				)
			) AS json_data
		) BETWEEN %d AND %d 
		AND DATE_FORMAT(createdAt, '%%Y-%%m-%%d') BETWEEN '%s' AND '%s'`,
		request.MinCount, request.MaxCount, request.StartDate, request.EndDate)

	rows, err := m.Conn.QueryContext(ctx, query)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.Record, 0)
	for rows.Next() {
		t := domain.Record{}
		marks := ""

		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.CreatedAt,
			&marks,
		)

		t.Marks = repository.StringToInts(marks)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}
