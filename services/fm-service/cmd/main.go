// services/fm-service/cmd/main.go
package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/sithuhlaing/erp-system/shared/utils"
)

func main() {
    serviceName := getEnv("SERVICE_NAME", "fm")
    port := getEnv("PORT", "8001")
    
    utils.InitLogger(serviceName)
    utils.Logger.Info("Starting Financial Management Service")

    r := gin.Default()

    // Health check
    r.GET("/health", func(c *gin.Context) {
        utils.SuccessResponse(c, "FM Service is healthy", gin.H{
            "service": serviceName,
            "version": "1.0.0",
        })
    })

    // Hello World endpoints
    r.GET("/", func(c *gin.Context) {
        utils.SuccessResponse(c, "Hello from Financial Management Service!", gin.H{
            "service": "fm",
            "domain": "Financial Management",
        })
    })

    r.GET("/api/v1/fm/hello", func(c *gin.Context) {
        utils.SuccessResponse(c, "Hello from FM API!", gin.H{
            "endpoints": []string{
                "GET /api/v1/fm/gl - General Ledger",
                "GET /api/v1/fm/ap - Accounts Payable", 
                "GET /api/v1/fm/ar - Accounts Receivable",
                "GET /api/v1/fm/reports - Financial Reports",
            },
        })
    })

    r.GET("/api/v1/fm/gl", func(c *gin.Context) {
        utils.SuccessResponse(c, "Hello from General Ledger!", gin.H{
            "module": "General Ledger",
            "features": []string{"Chart of Accounts", "Journal Entries", "Trial Balance"},
        })
    })

    r.GET("/api/v1/fm/ap", func(c *gin.Context) {
        utils.SuccessResponse(c, "Hello from Accounts Payable!", gin.H{
            "module": "Accounts Payable",
            "features": []string{"Vendor Invoices", "Payments", "Aging Reports"},
        })
    })

    r.GET("/api/v1/fm/ar", func(c *gin.Context) {
        utils.SuccessResponse(c, "Hello from Accounts Receivable!", gin.H{
            "module": "Accounts Receivable",
            "features": []string{"Customer Invoices", "Collections", "Credit Management"},
        })
    })

    r.GET("/api/v1/fm/reports", func(c *gin.Context) {
        utils.SuccessResponse(c, "Hello from Financial Reports!", gin.H{
            "module": "Financial Reports",
            "features": []string{"Balance Sheet", "Income Statement", "Cash Flow"},
        })
    })

    utils.Logger.WithField("port", port).Info("FM service starting")
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}