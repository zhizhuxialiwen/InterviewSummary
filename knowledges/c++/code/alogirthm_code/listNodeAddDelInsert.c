# include <stdio.h>
#include <malloc.h>

#define LEN sizeof(NODE)

typedef struct _NODE
{
    int val;
    struct _NODE *next;
} NODE, PNODE; 

//1. 打印所有节点
void print(PNODE *head)
{
    while(head) 
    {
        printf
        ("%3d", head->val);
        head = head->next;
    }
    printf("\n");
}

//2. 插入
//2.1 头插入
/*
链表分为空链表和非空链表的情况；
空链表的第一个参数为双指针**pHead，新插入的节点就是头指针，由于导致改动头指针，
因此必须把头指针设置为指针的指针，否则pHead仍然是一个空指针。
*/
void insertHead(PNODE ** pHead, int value) 
{
    PNODE *newNode = (PNODE *)malloc(LEN);
    newNode->val = value;
    newNode->next = *pHead;
    *pHead = newNode;
}

//2.2 尾部插入
void insertTail(PNODE ** pHead, int value)
{
    PNODE *tempNode = *pHead;
    PNODE *newNode = (PNODE *)malloc(LEN);
    newNode->val = value;
    newNode->next = NULL;
    if(*pHead == NULL){ //空链表
        newNode->next = *pHead;
        *pHead = newNode;
    } else {
        while(tempNode->next) {
            tempNode = tempNode->next;
        }
        tempNode->next = newNode;
        
    }
}

//3. 删除节点
//3.1 删除头结点
void delHead(PNODE **pHead) {
    if(*pHead == NULL){
        return;
    } else {
        PNODE *tempNode = *pHead;
        *pHead = (*pHead)->next;
        free(tempNode);
    }
}
//3.2删除尾结点
void delTail(PNODE **pHead) {
    PNODE *tempNode = *pHead;
    if(*tempNode == NULL) {
        return;
    } else if(tempNode->next == NULL) {
        free(tempNode);
        tempNode = NULL;
    }
}

//
//4.根据值查找节点
PNODE * findByVal(PNODE *head, int value)
{
    while(head != NULL && head->val != value) {
        head = head->next;
    }
    return head;
}

//4.1 根据值删除节点

void delByVal(PNODE **pHead, int value) 
{
    if(*pHead == NULL){
        return;
    } else {
        if((*pHead)->val == value) {
            deleteHead(pHead);
        } else {
            PNODE *tempNode = *pHead;
            while(tempNode->next != NULL && tempNode->next->val != value){
                tempNode = tempNode->next;
            }
            if(tempNode->next) {
                tempNode->next = tempNode->next->next;
                free(tempNode->next);
            }
      
        }
    }
}
//5. 根据索引查找结点
PNODE * findByIndex(PNODE *head, int index)
{
    if(index == 1) {
        return head;
    } else{
        int count = 1;
        while(head != NULL && index != count) {
            head = head->next;
            count++;
        }
    }
    return head;
} 

//5.1 根据索引插入节点
void insertByIndex(PNODE ** pHead, int index, int value) 
{
    if(index == 1) {
        insertHead(pHead, value);
    } else {
        PNODE *tempNode = findByIndex(*pHead, index - 1);
        if(tempNode == NULL) {
            return;
        } else {
            PNODE *newNode = (PNODE *)malloc(LEN);
            newNode->val = value;
            newNode->next = tempNode->next;
            tempNode->next = newNode;
        }
    }
}

//5.2 根据索引删除节点
void delByIndex(PNODE ** pHead, int index, int value) 
{
    if(index == 1) {
        delHead(pHead, value);
    } else {
        PNODE *tempNode = findByIndex(*pHead, index - 1);
        if(tempNode == NULL) {
            return;
        } else {
           tempNode->next = tempNode->next->next;
           free(tempNode->next);
        }
    }

    
void main()
{
    PNODE head = NULL;

    insertTail(&head,1);
    deleteHead(&head);
    insertTail(&head,2);
    insertTail(&head,3);
    insertTail(&head,4);
    insertTail(&head,5);
    insertTail(&head,6);

    print(head);
    insertByIndex(&head, 6, 9);
    print(head);
    //deleteByIndex(&head,3);
    deleteByVal(&head, 2);
    print(head);
    clear(&head);
    print(head);
    insertByIndex(&head,1,12);
    print(head);
}
}