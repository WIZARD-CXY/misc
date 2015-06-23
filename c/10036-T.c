#include<stdio.h>
#include<stdlib.h>
#include<string.h>
struct stud_node{
 int num;
 struct stud_node *next;
};
void Ptrint_Stu_Doc(struct stud_node *head);
void main()
{
 struct stud_node *L,*tail1,*tail2,*p1,*p2,*NEW;
 int num;
 int size=sizeof(struct stud_node);
 L=tail1=NULL;
 scanf("%d",&num);
 while(num!=-1){ 
 p1=(struct stud_node *)malloc(size);
 p1->num=num;
 p1->next=NULL;
 if(L==NULL)
 L=p1;
 else
 tail1->next=p1;
 tail1=p1;
 scanf("%d",&num);
 }
 NEW=tail2=NULL;
 p2=L;
 while (p2!=NULL) {
 if ( p2->num%2) {
 p1=(struct stud_node *)malloc(size);
 p1->num=p2->num;
 p1->next=NULL;
 if(NEW==NULL)
 NEW=p1;
 else
 tail2->next=p1;
 tail2=p1;
 }
 p2=p2->next;
 }

 Ptrint_Stu_Doc(NEW);
 free(L);
 free(NEW);
}
void Ptrint_Stu_Doc(struct stud_node *head)
{
 struct stud_node *ptr;
 if(head==NULL){
 printf("No Records\n");
 return;
 }
 for(ptr=head;ptr;ptr=ptr->next)
 printf("%d ",ptr->num);
 printf("\n");
} 