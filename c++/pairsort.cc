#include <algorithm>
#include <iostream>
#include <queue>
using namespace std;

typedef std::pair<size_t, double> PQNode;
int main() {
  auto cmp = [](const std::pair<size_t, double> &left, const std::pair<size_t, double>& right) {
    return left.second >= right.second;
  };

  std::priority_queue<PQNode, std::vector<PQNode>, decltype(cmp)> open_pq(cmp);
  open_pq.push(std::make_pair(0, 0.0));
  open_pq.push(std::make_pair(1, 1.0));
  open_pq.push(std::make_pair(2, 2.0));
 
  while(!open_pq.empty()) {
   size_t curidx = open_pq.top().first; 
   open_pq.pop();
   cout<< curidx <<endl;
  }
}
   
