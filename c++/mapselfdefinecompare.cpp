class person1
{  
public:  
    string name;  
    int age; 
    person1(string s,int i):name(s),age(i){}  
	//map必须重载键类型的<符号
	const bool operator<(const person1 &a)const
	{
		if(age!=a.age)
			return age<a.age;
		else
			return name<a.name;
	}
};  
  
class person2
{  
public:  
    string name;  
    int age; 
    person2(string s,int i):name(s),age(i){}  

	//unordered_map必须重载键类型的==符号，才能在哈希值相等的情况下找到准确的键值对
	const bool operator==(const person2 &a)const
	{
		return age==a.age && name==a.name;
	}
}; 

//person2类的哈希函数
struct person2_hash
{
	//unordered_map必须定义键类型的哈希函数
	size_t operator()(const person2 &p)const  
    {  
        return hash<int>()(p.age)+hash<char>()(p.name[0]);  
    }  
};

//测试代码
//map操作
map<person1,string> map_person;
person1 p1("Tom",15);
person1 p2("Tim",22);
//map_person.emplace("Tim",22,"driver"); 该行语法错误
map_person[p1]="student";
map_person[p2]="driver";
cout<<map_person[p1]<<endl;//输出 student
cout<<map_person[p2]<<endl;//输出 driver

//unordered_map操作
unordered_map<person2,string,person2_hash> unordered_person;
person2 x1("Ton",14);//x1和x2的哈希值相等
person2 x2("Tim",14);
person2 x3("John",50);
unordered_person[x1]="I am Ton";
unordered_person[x2]="I am Tim";
unordered_person[x3]="I am John";
cout<<unordered_person[x1]<<endl;//输出  I am Ton
cout<<unordered_person[x2]<<endl;//输出  I am Tim
cout<<unordered_person[x3]<<endl;//输出  I am John