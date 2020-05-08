// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	"github.com/byuoitav/common"
// 	"github.com/byuoitav/common/log"
// 	"github.com/byuoitav/common/v2/auth"
// 	"github.com/byuoitav/symetrix-driver/symetrix"
// 	"github.com/labstack/echo"
// )

// func main() {
// 	log.SetLevel("info")
// 	port := ":8111"
// 	router := common.NewRouter()

// 	// Functionality Endpoints
// 	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))
// 	write.GET("/:address/block/:block/volume/:level", func(ctx echo.Context) error {
// 		dsp := symetrix.DSP{}
// 		address := ctx.Param("address")
// 		block := ctx.Param("block")
// 		gainBlock, _, err := parseBlock(block)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		level := ctx.Param("level")
// 		volume, err := strconv.Atoi(level)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		dsp.Address = address
// 		c := ctx.Request().Context()
// 		err = dsp.SetVolumeByBlock(c, gainBlock, volume)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		return ctx.String(http.StatusOK, "success")
// 	})
// 	write.GET("/:address/block/:block/muted/true", func(ctx echo.Context) error {
// 		dsp := symetrix.DSP{}
// 		address := ctx.Param("address")
// 		block := ctx.Param("block")
// 		_, muteBlock, err := parseBlock(block)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		dsp.Address = address
// 		c := ctx.Request().Context()
// 		err = dsp.SetMutedByBlock(c, muteBlock, true)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		return ctx.String(http.StatusOK, "success")
// 	})
// 	write.GET("/:address/block/:block/muted/false", func(ctx echo.Context) error {
// 		dsp := symetrix.DSP{}
// 		address := ctx.Param("address")
// 		block := ctx.Param("block")
// 		_, muteBlock, err := parseBlock(block)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		dsp.Address = address
// 		c := ctx.Request().Context()
// 		err = dsp.SetMutedByBlock(c, muteBlock, false)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		return ctx.String(http.StatusOK, "success")
// 	})

// 	read := router.Group("", auth.AuthorizeRequest("read-state", "room", auth.LookupResourceFromAddress))

// 	read.GET("/:address/block/:block/volume", func(ctx echo.Context) error {
// 		dsp := symetrix.DSP{}
// 		address := ctx.Param("address")
// 		block := ctx.Param("block")
// 		gainBlock, _, err := parseBlock(block)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		dsp.Address = address
// 		c := ctx.Request().Context()
// 		volume, err := dsp.GetVolumeByBlock(c, gainBlock)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		return ctx.String(http.StatusOK, fmt.Sprintf("%s: %v", gainBlock, volume))

// 	})
// 	read.GET("/:address/block/:block/muted", func(ctx echo.Context) error {
// 		dsp := symetrix.DSP{}
// 		address := ctx.Param("address")
// 		block := ctx.Param("block")
// 		_, muteBlock, err := parseBlock(block)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		dsp.Address = address
// 		c := ctx.Request().Context()
// 		muted, err := dsp.GetMutedByBlock(c, muteBlock)
// 		if err != nil {
// 			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
// 		}

// 		return ctx.String(http.StatusOK, fmt.Sprintf("%s: %v", muteBlock, muted))
// 	})

// 	// log level endpoints
// 	router.PUT("/log-level/:level", log.SetLogLevel)
// 	router.GET("/log-level", log.GetLogLevel)

// 	server := http.Server{
// 		Addr:           port,
// 		MaxHeaderBytes: 1024 * 10,
// 	}

// 	router.StartServer(&server)
// }

// func parseBlock(block string) (string, string, error) {
// 	parsedBlock := strings.Split(block, "|")
// 	if len(parsedBlock) == 1 {
// 		return "", "", fmt.Errorf("block is not in the correct format. Expecting gain#|mute#, recieved: %s", block)
// 	}

// 	gainBlock := parsedBlock[0]
// 	muteBlock := parsedBlock[1]

// 	return gainBlock, muteBlock, nil
// }
