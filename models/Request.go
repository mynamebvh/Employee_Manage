package models

import (
	db "employee_manage/config"
	modelErrors "employee_manage/constant"
	"employee_manage/models/dto"
	"employee_manage/utils"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Request struct {
	ID         int       `gorm:"primaryKey" json:"id" `
	Type       string    `json:"type"`
	Content    string    `json:"content"`
	Status     string    `json:"status"`
	UserID     int       `gorm:"not null;column:user_id" json:"user_id"`
	User       User      `json:"-"`
	ApprovedBy *int      `gorm:"column:approved_by" json:"approved_by"`
	Manager    User      `gorm:"foreignKey:ApprovedBy" json:"-"`
	CreatedAt  time.Time `json:"create_at" gorm:"autoCreateTime:mili"`
	UpdatedAt  time.Time `json:"-" gorm:"autoUpdateTime:mili"`
}

func CreateRequest(re *Request) (err error) {
	err = db.DB.Create(re).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError modelErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return err
		}
		switch newError.Number {
		case 1062:
			err = modelErrors.NewAppErrorWithType(modelErrors.ResourceAlreadyExists)
			return

		default:
			err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}

	return
}

func GetRequests(c *gin.Context, role map[string]interface{}) (data dto.QueryPagination, err error) {
	var page, limit int
	var count int64
	var result []dto.QueryGetRequest

	if role["name"] == "admin" {
		if err = db.DB.
			Table("requests as rq").
			Select("rq.type", "rq.content", "rq.status", "u.full_name as full_name", "u1.full_name as approved_by").
			Joins("left join users as u on rq.user_id = u.id").
			Joins("left join users as u1 on rq.approved_by = u1.id").
			Scopes(utils.Paginate(c, &page, &limit)).
			Scan(&result).Error; err != nil {
			return
		}

		if err = db.DB.Table("requests").Count(&count).Error; err != nil {
			return
		}

		data = dto.QueryPagination{
			Current:  page,
			Total:    count,
			PageSize: limit,
			Data:     result,
		}
	} else {
		departmentID := role["department_id"].(int)
		subQuery := db.DB.Table("users").Select("id").Where("department_id = ?", departmentID)

		db.DB.Table("requests as rq").Where("user_id in (?)", subQuery).
			Select("rq.type", "rq.content", "rq.status", "u.full_name as full_name", "u1.full_name as approved_by").
			Joins("left join users as u on rq.user_id = u.id").
			Joins("left join users as u1 on rq.approved_by = u1.id").
			Scopes(utils.Paginate(c, &page, &limit)).
			Scan(&result)

		if err = db.DB.Table("requests").Where("user_id in (?)", subQuery).Count(&count).Error; err != nil {
			return
		}

		data = dto.QueryPagination{
			Current:  page,
			Total:    count,
			PageSize: limit,
			Data:     result,
		}
	}
	return
}

func GetRequestByID(id int) (re dto.QueryGetRequest, err error) {
	err = db.DB.Table("requests as rq").
		Select("rq.type", "rq.content", "u.full_name as full_name", "a.full_name as approved_by", "rq.status").
		Where("rq.id = ?", id).
		Joins("left join users as u on rq.user_id = u.id").
		Joins("left join users as a on rq.approved_by = a.id").
		Scan(&re).
		Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
		default:
			err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}

	return
}

func UpdateRequestByID(id int, requestMap map[string]interface{}) (re Request, err error) {
	re.ID = id
	err = db.DB.Model(&re).
		Select("status", "approved_by").
		Updates(requestMap).
		Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError modelErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}

		switch newError.Number {
		case 1062:
			err = modelErrors.NewAppErrorWithType(modelErrors.ResourceAlreadyExists)
			return

		default:
			err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}

	err = db.DB.Where("id = ?", id).First(&re).Error

	if err != nil {
		err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
		return
	}

	return
}
