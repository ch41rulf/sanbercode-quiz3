package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

func SegitigaSamaSisiHandler(c *gin.Context) {
	hitung := c.Query("hitung")
	alasStr := c.Query("alas")
	tinggiStr := c.Query("tinggi")

	alas, err := strconv.ParseFloat(alasStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah 'alas'"})
		return
	}
	tinggi, err := strconv.ParseFloat(tinggiStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah'tinggi'"})
		return
	}

	var hasil float64
	switch hitung {
	case "luas":
		hasil = 0.5 * alas * tinggi
	case "keliling":
		hasil = 3 * alas
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah for 'hitung'"})
		return
	}

	c.JSON(200, gin.H{"result": hasil})
}

func PersegiHandler(c *gin.Context) {
	hitung := c.Query("hitung")
	sisiStr := c.Query("sisi")

	sisi, err := strconv.ParseFloat(sisiStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah 'sisi'"})
		return
	}

	var hasil float64
	switch hitung {
	case "luas":
		hasil = sisi * sisi

	case "keliling":
		hasil = 4 * sisi
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah for 'hitung'"})
	}
	c.JSON(200, gin.H{"result": hasil})
}

func PersegiPanjangHandler(c *gin.Context) {
	hitung := c.Query("hitung")
	panjangStr := c.Query("panjang")
	lebarStr := c.Query("lebar")

	panjang, err := strconv.ParseFloat(panjangStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah 'panjang'"})
		return
	}
	lebar, err := strconv.ParseFloat(lebarStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah 'lebar'"})
		return
	}

	var hasil float64
	switch hitung {
	case "luas":
		hasil = panjang * lebar

	case "keliling":
		hasil = (2 * panjang) + (2 * lebar)
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah for 'hitung'"})
	}
	c.JSON(200, gin.H{"result": hasil})
}

func LingkaranHandler(c *gin.Context) {
	hitung := c.Query("hitung")
	jariJariStr := c.Query("jariJari")

	jariJari, err := strconv.ParseFloat(jariJariStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah 'jariJari'"})
		return
	}

	var hasil float64
	switch hitung {
	case "luas":
		hasil = math.Pi * jariJari * jariJari

	case "keliling":
		hasil = 2 * math.Pi * jariJari
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": "nilai salah for 'hitung'"})
	}
	c.JSON(200, gin.H{"result": hasil})
}
