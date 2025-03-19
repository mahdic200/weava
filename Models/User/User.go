package User

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Utils/Http"
	"gorm.io/gorm"
)

var fillable = []string{
    "first_name",
    "last_name",
    "email",
    "phone",
    "image",
    "password",
};

func Create(tx *gorm.DB, args map[string]string) *gorm.DB {
    fields := []string{}
    question_marks := []string{}
    validated_args := []any{}
    for key, value := range args {
        if slices.Contains(fillable, key) {
            fields = append(fields, key)
            question_marks = append(question_marks, "?")
            validated_args = append(validated_args, value)
        }
    }
    fields = append(fields, "created_at")
    question_marks = append(question_marks, "?")
    validated_args = append(validated_args, time.Now())

    query := fmt.Sprintf("INSERT INTO users (%s) VALUES (%s)", strings.Join(fields, ", "), strings.Join(question_marks, ", "))
    tx = tx.Exec(query, validated_args...)

    return tx
}

func Paginate(tx *gorm.DB, c *fiber.Ctx) (*gorm.DB, Http.PaginationMetadata) {
    return Http.Paginate(tx, c, 15)
}
