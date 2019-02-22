package main

import (
	"./dbcon"
	"./gueztli"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"runtime"
	"strconv"

	"os"
	"sync"
)

func gupho() {
	//var ret *int
	scanDir("E:\\mzitu", 0)
	//var wg sync.WaitGroup
	//var count, numCpu = 0, runtime.NumCPU()
	//list := make([]os.FileInfo, numCpu)
	//iDir, oDir := "E:\\mzi\\tuku_1", "E:\\mzi\\tuku"
	//dir, err := ioutil.ReadDir(iDir)
	//
	//if err != nil {
	//	//fmt.Println("WWW")
	//}
	//for _, fi := range dir {
	//	println(count)
	//	if fi.IsDir() { // 目录, 递归遍历
	//		//dbcon.Insert("Path", strconv.Itoa(count), fi.Name(), thumbPath, )
	//		//err := os.Rename(dirPth+PthSep+fi.Name(), newPath)
	//		//checkErr(err)
	//		//GetFilesAndDirs(newPath)
	//	}
	//
	//	//fmt.Println(fi.Name())
	//	list[count] = fi
	//	count++
	//	if math.Mod(float64(count), float64(numCpu)) == 0 {
	//		guetzliEncode4(iDir, oDir, &wg, list)
	//		count = 0
	//	}
	//
	//}

}

func scanDir(iDir string, count int) {
	dir, err := ioutil.ReadDir(iDir)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	numCPU := runtime.NumCPU()
	list := make([]os.FileInfo, numCPU)
	oDir := "E:\\mzigu\\gku_"

	for _, di := range dir {
		gCount := 0
		finC := 0
		if di.IsDir() { // 目录, 递归遍历
			fmt.Println(di.Name())
			count++
			path := dbcon.SearchPath(di.Name())
			if path != nil {
				continue
			}

			//scanDir(iDir + "\\" + fi.Name(), count)
			dbcon.Insert("Path", strconv.Itoa(count), di.Name(), "tuku_"+strconv.Itoa(count))
			files, err := ioutil.ReadDir(iDir + "\\" + di.Name())
			if err != nil {
				panic(err)
			}

			//fmt.Println(len(files))
			for _, fi := range files {
				_, err := os.Stat(oDir + strconv.Itoa(count) + "\\" + fi.Name())

				if err != nil {
					if os.IsNotExist(err) {
						list[gCount] = fi
						gCount++
						finC++
						if math.Mod(float64(gCount), float64(numCPU)) == 0 {
							guetzliEncode4(iDir+"\\"+di.Name(), oDir+strconv.Itoa(count), &wg, list)
							gCount = 0
						} else if finC == len(files) {
							fmt.Println(list)
							fmt.Println(finC)
							mod := int(math.Mod(float64(finC), float64(numCPU)))
							guetzliEncode4(iDir+"\\"+di.Name(), oDir+strconv.Itoa(count), &wg, list[:mod])
						}
					}
				}
			}

		}
	}
}

func guetzliEncode4(iDir string, oDir string, wg *sync.WaitGroup, list []os.FileInfo) {
	PthSep := string(os.PathSeparator)
	wg.Add(len(list))
	for _, val := range list {
		fmt.Println(val.Name())
		go func(val os.FileInfo) {
			defer wg.Done()
			//fmt.Println(iDir + PthSep + val.Name())
			m0 := gueztli.LoadImage(iDir + PthSep + val.Name())
			err := os.MkdirAll(oDir, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println("guetzliEncode")
			data2 := gueztli.GuetzliEncode(m0, 84)

			//fmt.Println(newPath + PthSep + val.Name())
			//fmt.Println("WriteFile::::::::::" + newPath + PthSep + val.Name())
			if err := ioutil.WriteFile(oDir+PthSep+val.Name(), data2, 0666); err != nil {
				log.Println(err)
			}
		}(val)
	}
	wg.Wait()
}
