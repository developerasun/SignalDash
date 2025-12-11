package service

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/developerasun/SignalDash/server/config"
	"github.com/developerasun/SignalDash/server/dto"
	"github.com/developerasun/SignalDash/server/models"
	"github.com/developerasun/SignalDash/server/sderror"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func cleanup(t *testing.T) *gorm.DB {
	absPath, err := filepath.Abs("../config")
	require.NoError(t, err)

	env := config.NewEnvironment(absPath, "options").Instance
	testDB := env.GetString("server.database.test")

	dbPath, err := filepath.Abs("../" + testDB)
	require.NoError(t, err)

	_ = os.Remove(dbPath)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&models.Indicator{})
	require.NoError(t, err)

	return db
}

func TestFindAndCreateDollarIndex(t *testing.T) {
	t.Skip()
	db := cleanup(t)
	_, fErr := FindLatestDollarIndex(db)

	// @dev should be clean
	require.ErrorIs(t, fErr, sderror.ErrEmptyStorage)

	dxy := "89.44"
	cErr := CreateDollarIndex(db, dxy)
	require.NoError(t, cErr)

	_, rfErr := FindLatestDollarIndex(db)
	require.NotErrorIs(t, rfErr, sderror.ErrEmptyStorage)
}

func TestNewHttp(t *testing.T) {
	res, err := DoHttpGet([]string{"https://api.bithumb.com/v1/ticker?markets=KRW-USDT"})

	require.NoError(t, err)
	require.True(t, len(res) == 1)
	require.True(t, strings.Contains(res[0], "KRW-USDT"))

	var bithumbApiResType []dto.BithumbApiItem
	uErr := json.Unmarshal([]byte(res[0]), &bithumbApiResType)
	t.Log("value: ", bithumbApiResType[0].OpeningPrice)
	require.NoError(t, uErr)
}
