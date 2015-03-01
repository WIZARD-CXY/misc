/*************************************************************************
	> File Name: list.cpp
	> Author: 
	> Mail: 
	> Created Time: Sat 28 Feb 2015 11:03:52 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;

struct ListNode{
    int val;
    ListNode *next;
};

ListNode* CreateNode(int value){
    ListNode *node = new ListNode();
    node->val=value;
    node->next=NULL;

    return node;
}

void PrintList(ListNode *head){
    
    while(head){
        cout<<head->val<<" ";
        head=head->next;
    }
    cout<<endl;

}

void AddToTail(ListNode* &pHead,int value){
    ListNode *node = CreateNode(value);

    if(pHead==NULL){
        pHead = node;
    }else{
        ListNode *pNode=pHead;

        while(pNode->next){
            pNode=pNode->next;
        }
        pNode->next=node;
    }
}

ListNode* ReverseList(ListNode* head){
    if(head==NULL){
        return NULL;
    }
    ListNode *node = head; //current node
    ListNode *pre = NULL; //current node's previous node
    ListNode *ReversedHead=NULL;

    while(node){
        ListNode *next = node->next;//record the node's next node
        
        if(next==NULL){
            //node is the tail
            ReversedHead=node;
        }

        node->next=pre;
        pre=node;
        node=next;
    }
    return ReversedHead;
}

int main(){
    ListNode* head=CreateNode(1);

    AddToTail(head,2);
    /*AddToTail(head,3);
    AddToTail(head,4);
    AddToTail(head,5);
    AddToTail(head,6);
    AddToTail(head,7);
    AddToTail(head,8);
*/
    PrintList(head);

    ListNode* reversedHead=ReverseList(head);

    PrintList(reversedHead);

}
