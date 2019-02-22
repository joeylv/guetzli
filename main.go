package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"

	"./dbcon"
	"./gueztli"
)

// Conf p
type Conf struct {
	iDir string
	oDir string
	fi   string
}

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
//jobs <-chan int：只能接收数据
//　　results chan<- int：只能发送数据
func worker(id int, jobs <-chan Conf, results chan<- Conf) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		//time.Sleep(time.Second)
		//fmt.Println(j.iDir)
		//fmt.Println(j.oDir)
		//fmt.Println(j.fi)
		_, err := os.Stat(j.oDir + "\\" + j.fi)

		if err != nil {
			if os.IsNotExist(err) {
				m0 := gueztli.LoadImage(j.iDir + "\\" + j.fi)
				err := os.MkdirAll(j.oDir, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
				data2 := gueztli.GuetzliEncode(m0, 84)
				if err := ioutil.WriteFile(j.oDir+"\\"+j.fi, data2, 0666); err != nil {
					log.Println(err)
				}
			}
		}
		fmt.Println("worker", id, "finished job", j)
		results <- j
	}
}

func main() {
	//Rname("E:/mzitu")
	//RCount("E:/mzitu")

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	list := getFiles("E:/mzitu")
	//fmt.Println(len(list))
	jobs := make(chan Conf, len(list))
	results := make(chan Conf, len(list))
	//getFiles("E:/mzitu",list)
	//for l := range list {
	//	fmt.Println(<-l)
	//}
	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= runtime.NumCPU(); w++ {
		go worker(w, jobs, results)
	}

	//fmt.Println(cap(jobs))
	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	//var iDir = "E:/mzitu"
	//var oDir = "E:/mzigun/gku_"
	for j := 0; j < cap(jobs); j++ {
		jobs <- list[j]
	}
	//}
	close(jobs)

	// Finally we collect all the results of the work.
	for a := 0; a <= cap(jobs); a++ {
		<-results
	}
}

func getFiles(iDir string) []Conf {
	dir, err := ioutil.ReadDir(iDir)
	if err != nil {
		panic(err)
	}
	var oDir = "E:/mzigu/gku_"
	var list []Conf
	var count = 0
	PthSep := string(os.PathSeparator)
	for _, di := range dir {
		if di.IsDir() { // 目录, 递归遍历
			path := dbcon.SearchPath(di.Name())
			if path != nil {
				//fmt.Println(path)
				count, _ = strconv.Atoi(path.Name)
				//continue
			} else {
				count++
			}

			fmt.Println(count)
			dbcon.Insert("Path", strconv.Itoa(count), di.Name(), "tuku_"+strconv.Itoa(count))
			//scanDir(iDir + "\\" + fi.Name(), count)
			fmt.Println(di.Name())
			files, err := ioutil.ReadDir(iDir + PthSep + di.Name())
			if err != nil {
				panic(err)
			}
			//fmt.Println(len(files))
			for _, fi := range files {
				_, err := os.Stat(oDir + strconv.Itoa(count))
				if err != nil {
					if os.IsNotExist(err) {
						//fmt.Println(count)
						//list[count] = Conf{iDir + PthSep + di.Name(), oDir + strconv.Itoa(count), fi.Name()}
						list = append(list, Conf{iDir + PthSep + di.Name(), oDir + strconv.Itoa(count), fi.Name()})
						//fmt.Println(list)
					}
				}
			}
		}
	}
	return list
}

//RCount R
func RCount(iDir string) {
	dir, err := ioutil.ReadDir(iDir)
	if err != nil {
		panic(err)
	}
	var oDir = "E:/mzigu/gku_"
	var count = 0
	for _, di := range dir {
		if di.IsDir() { // 目录, 递归遍历
			count++

			fmt.Println(count)

			up := dbcon.Insert("Path", strconv.Itoa(count), di.Name(), "tuku_"+strconv.Itoa(count))
			if up > 0 {
				os.Rename("E:/mzigu/"+di.Name(), oDir+strconv.Itoa(count))
			}

		}
	}
}

//Rname R
func Rname(iDir string) {
	dir, err := ioutil.ReadDir(iDir)
	if err != nil {
		panic(err)
	}
	for _, di := range dir {
		if di.IsDir() { // 目录, 递归遍历
			path := dbcon.SearchPath(di.Name())
			if path != nil {
				//count, _ = strconv.Atoi(path.Name)
				os.Rename("E:/mzigu/gku_"+path.Name, "E:/mzigu/"+path.Path)
				//fmt.Println(path)

				//continue
			}
		}
	}
}
