package main
import (
	"bufio"
	"fmt"
	"os"
	"io"
)

func fileOpen1()  {
	fileObj, err := os.Open("./main.go")	
	if err != nil {
		fmt.Printf("open file err: %v", err)
		return 
	}
	defer fileObj.Close()
	var tem = make([]byte, 127)
	for {
		n , err := fileObj.Read(tem)
		if err == io.EOF {
			fmt.Println("文件已读完")
			return
		}
		if err != nil {
			fmt.Printf("read flie err : %v", err)
			return
		}
		fmt.Println(string(tem[:n]))
	}

}

func fileRead()  {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file err: %v", err)
		return 
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件已读完")
			return
		}
		if err != nil {
			fmt.Printf("read flie err : %v", err)
			return
		}
		fmt.Println(line)
	}
	
}

// func useIoutil()  {
// 	file, err := ioutil.ReadFile("./main.go")
// 	if err != nil {
// 		fmt.Printf("ioutil read flie err : %v", err)
// 		return
// 	}
// 	fmt.Println(file)
// }

func main()  {
	// fileOpen1()
	// fileRead()
	// useIoutil()
}