/*
* @Author: sottxiong
* @Date:   2019-12-25 06:39:19
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-12-25 07:39:42
*/
package engine

import(
	"fmt"
	"os"
	"strings"
	"github.com/scott-x/gutils/cmd"
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/gutils/cl"
)

func Run(){
	cmd.AddQuestion("path", "The iconfont folder path: ", "Please drag the folder here: ", "^/.*font_.*")
	answers := cmd.Exec()
	path := strings.Trim(answers["path"]," ")
	flag :=runTask(path)
	if flag{
		cl.BoldCyan.Printf("Done, please copy folder iconfont into your React project...")
	}else{
		cl.BoldRed.Printf("The file is incompleted, please double check, then try agagin...")
	}
	
}

func runTask(path string) bool{
	fs.CreateDirIfNotExist("./iconfont");
	var flag bool=true
	if fs.IsExist(path+"/iconfont.css") {
		fs.Copy(path+"/iconfont.css", "./iconfont/iconfont.css");
	}else {
		flag = false
		fmt.Printf("not found: iconfont.css \n")
	}
	if fs.IsExist(path+"/iconfont.eot"){
		fs.Copy(path+"/iconfont.eot", "./iconfont/iconfont.eot");
	}else{
		flag = false
		fmt.Printf("not found: iconfont.eot \n")
	}
	if fs.IsExist(path+"/iconfont.svg"){
		fs.Copy(path+"/iconfont.svg", "./iconfont/iconfont.svg");
	}else{
		flag = false
		fmt.Printf("not found: iconfont.svg \n")
	}
	if fs.IsExist(path+"/iconfont.ttf"){
		fs.Copy(path+"/iconfont.ttf", "./iconfont/iconfont.ttf");
	}else{
		flag = false
		fmt.Printf("not found: iconfont.ttf \n")
	}
	if fs.IsExist(path+"/iconfont.woff"){
		fs.Copy(path+"/iconfont.woff", "./iconfont/iconfont.woff");
	}else{
		flag = false
		fmt.Printf("not found: iconfont.woff \n")
	}
	if fs.IsExist(path+"/iconfont.woff2"){
		fs.Copy(path+"/iconfont.woff2", "./iconfont/iconfont.woff2");
	}else{
		flag = false
		fmt.Printf("not found: iconfont.woff2 \n")
	}
	
	content, err :=fs.ReadFile1("./iconfont/iconfont.css");
	if err!=nil{
		panic(err)
	}
	var newContent []string
	newContent = append(newContent,"import { createGlobalStyle } from 'styled-components';",)
	newContent = append(newContent,"")
	newContent = append(newContent,"export const GlobalIcon = createGlobalStyle`")
	arr := strings.Split(content,"\n")
	for x, c := range arr{
		if x>15 {
			break
		}
		if strings.Contains(c,"url('iconfont") {
			c=strings.Replace(c,"url('iconfont","url('./iconfont",-1)
		}
		newContent = append(newContent,c)
		
	}
	newContent = append(newContent,"`")
	fs.WriteString("./iconfont/iconfont.js", strings.Join(newContent,"\n") )
	//delete iconfont.css
    err =os.Remove("./iconfont/iconfont.css")
    if err!=nil{
    	panic(err)
    }
	return flag
}