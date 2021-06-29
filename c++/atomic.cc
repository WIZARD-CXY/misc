std::atomic<int> ai;
ai++;
ai.store(100);
int a =ai.load();


void func(){
  std::shared_ptr<A> sp = std:make_shared<A>();
}



void vecfunc() {
   vector<int> vec;
   cout<<"vector size "<<vec.size()<<endl;
   cout<<"vector capacity "<<vec.capacity()<<endl;


   vec.push_back(1);
   vec.emplace_back(2);
   vec.push_back(3);

   cout<<endl;
   vec.reserve(200);

   cout<<endl;

   vec.resize(20);

   vec.clear()

   vector<int>().swap(vec);

}


void erase(vector<int> &vec, int a){
   for (auto iter = vec.begin();iter!=vec.end()){
     if (*iter == a){
        iter=vec.erase(iter);
     } else{
        iter++;
     }
}

int main(){
   std::map<int,string> mmap;
   map[1]="hello";
   cout<<map[1]<<endl;
} 
