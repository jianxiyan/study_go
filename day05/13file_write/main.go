package main
import (
	"fmt"
	"os"
	"bufio"
)

func writeFile01()  {
	file, err := os.OpenFile("xxx.txt", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 6)
	if err != nil {
		fmt.Printf("openfile is err : %v", err)
		return
	}
	defer file.Close()
	file.Write([]byte("sha he \n"))
	file.WriteString("laalaala")
}

func bufioWriteFile()  {
	file, err := os.OpenFile("xxx.txt", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 111)
	if err != nil {
		fmt.Printf("openfile is err : %v", err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	for i := 0; i<10; i++ {
		write.WriteString("我来啦\n")
	}
	write.Flush()
}

func main()  {
	bufioWriteFile()
}