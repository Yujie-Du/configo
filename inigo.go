package configo
import "io/ioutil"
import "os"
import _ "fmt"

type iniconfig struct{
	data map[string]map[string]string
	filename string
}
//NewIni:新建一个iniconfig对象。filename:ini文件名
func NewIni(filename string)(ic *iniconfig){
	file,err:=os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err!=nil{
		panic("err.Error")
	}
	defer file.Close()
	bstr,err:=ioutil.ReadAll(file)
	if err!=nil{
		panic("err.Error")
	}
	ic=new(iniconfig)
	ic.data=make(map[string]map[string]string)
	ic.filename=filename
	str:=[]rune(string(bstr))
	iniNext(str,ic.data)
	return
}
func(ini *iniconfig)Get(block,key string)(value string,ok bool){
	bl,ok:=ini.data[block]
	if !ok{
		return
	}
	value,ok=bl[key]
	return
}
func iniNext(str []rune,data map[string]map[string]string){
	block:="default"
	data[block]=make(map[string]string)
	for len(str)>0{
		idx:=index(str,rune('\n'))
		if name,ok:=checkBlockName(str[:idx]);ok{
			block=name
			data[block]=make(map[string]string)
		}else if key,value,ok:=checkKeyValue(str[:idx]);ok{
			data[block][key]=value
		}
		if idx==len(str){
			break
		}
		str=str[idx+1:]
	}
}
func checkBlockName(data []rune)(name string,ok bool){
	data=strip(data)
	if len(data)==0||data[0]!=rune('[')||data[len(data)-1]!=rune(']'){
		return
	}
	data=data[1:len(data)-1]
	name,ok=string(strip(data)),len(data)>0
	return
}
func checkKeyValue(data []rune)(key,value string,ok bool){
	idx:=index(data,rune('='))
	if idx==len(data){
		return
	}
	key,value=string(strip(data[:idx])),string(strip(data[idx+1:]))
	ok=len(key)>0&&len(value)>0
	return
}
func strip(data []rune)(data2 []rune){
	rib:=[]rune(" \n\t\r")
	for len(data)>0&&index(rib,data[0])<len(rib){
		data=data[1:]
	}
	for len(data)>0&&index(rib,data[len(data)-1])<len(rib){
		data=data[:len(data)-1]
	}
	return data
}
func index(data []rune,aim rune)(idx int){
	for idx=0;idx<len(data);idx+=1{
		if data[idx]==aim{
			return
		}
	}
	return
}