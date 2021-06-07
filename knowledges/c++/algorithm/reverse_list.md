# 单链表反转详解（4种算法实现）

[单链表反转网址](http://c.biancheng.net/view/8105.html)

通过前面章节的学习，读者已经对单链表以及它的用法有了一个完整的了解。在此基础上，本节再带领大家研究一个和单链表有关的问题，即如何实现单链表的反转。

![reverse_list1](../../../images/reverse_list1.PNG)

通过对比图 1 和 图 2 中的链表不难得知，所谓反转链表，就是将链表整体“反过来”，将头变成尾、尾变成头。那么，如何实现链表的反转呢？

常用的实现方案有 4 种，这里分别将它们称为迭代反转法、递归反转法、就地逆置法和头插法。值得一提的是，递归反转法更适用于反转不带头节点的链表；其它 3 种方法既能反转不带头节点的链表，也能反转带头节点的链表。

本节将以图 1 所示，即不带头节点的链表为例，给大家详细讲解各算法的实现思想。

## 1、迭代反转链表

该算法的实现思想：从当前链表的首元节点开始，一直遍历至链表的最后一个节点，这期间会逐个改变所遍历到的节点的指针域，另其指向前一个节点。

![reverse_list2](../../../images/reverse_list2.PNG)

![reverse_list3](../../../images/reverse_list3.PNG)

注意，这里只需改变 mid 所指节点的指向即可，不用修改 3 个指针的指向。

5) 最后只需改变 head 头指针的指向，另其和 mid 同向，就实现了链表的反转。

如下是实现整个过程的代码：
//迭代反转法，head 为无头节点链表的头指针

```c++
link * iteration_reverse(link* head) {
    if (head == NULL || head->next == NULL) {
        return head;
    }
    else {
        link * beg = NULL;
        link * mid = head;
        link * end = head->next;
        //一直遍历
        while (1)
        {
            //修改 mid 所指节点的指向
            mid->next = beg;
            //此时判断 end 是否为 NULL，如果成立则退出循环
            if (end == NULL) {
                break;
            }
            //整体向后移动 3 个指针
            beg = mid;
            mid = end;
            end = end->next;
        }
        //最后修改 head 头指针的指向
        head = mid;
        return head;
    }
}
```

修改后的代码：

```c++
link * iteration_reverse(link* head) {
    if (head == NULL || head->next == NULL) {
        return head;
    } 

    link * beg = NULL;
    link * cur = head;
    //一直遍历
    while (head->next != null)
    {
        //摘除当前节点，暂存后继节点
        link * end = head->next;
        //修改 curr 所指节点的指向
        cur->next = beg;
        //暂存当前节点
        beg = cur;
        //访问下一节点
        cur = end;
    }
    return cur;
    
}
```

## 2、递归反转链表

和迭代反转法的思想恰好相反，递归反转法的实现思想是从链表的尾节点开始，依次向前遍历，遍历过程依次改变各节点的指向，即另其指向前一个节点。

鉴于该方法的实现用到了递归算法，不易理解，因此和讲解其他实现方法不同，这里先给读者具体的实现代码，然后再给大家分析具体的实现过程：

```c++
link* recursive_reverse(link* head) {
    //递归的出口
    if (head == NULL || head->next == NULL)     // 空链或只有一个结点，直接返回头指针
    {
        return head; //倒数第二节点
    }
  
    //一直递归，找到链表中最后一个节点
    link *new_head = recursive_reverse(head->next);
    //当逐层退出时，new_head 的指向都不变，一直指向原链表中最后一个节点；
    //递归每退出一层，函数中 head 指针的指向都会发生改变，都指向上一个节点。
    //每退出一层，都需要改变 head->next 节点指针域的指向，同时令 head 所指节点的指针域为 NULL。
    //head->next为倒数第二个节点，head->next->next为最后一个节点
    head->next->next = head;
    head->next = NULL;
    //每一层递归结束，都要将新的头指针返回给上一层。由此，即可保证整个递归过程中，能够一直找得到新链表的表头。
    return new_head;
    
}
```

仍以图 1 中的链表为例，则整个递归实现反转的过程如下：

![reverse_list4](../../../images/reverse_list4.PNG)

![reverse_list5](../../../images/reverse_list5.PNG)

## 3、头插法反转链表

头插法：指在原有链表的基础上，依次将位于链表头部的节点摘下，然后采用从头部插入的方式生成一个新链表，则此链表即为原链表的反转版。

仍以图 1 所示的链表为例，接下来为大家演示头插反转法的具体实现过程：

![reverse_list6](../../../images/reverse_list6.PNG)

![reverse_list7](../../../images/reverse_list7.PNG)

由此，就实现了对原链表的反转，新反转链表的头指针为 new_head。

如下为以头插法实现链表反转的代码：

```c++
link * head_reverse(link * head) {
    link * new_head = NULL;
    link * temp = NULL;
    if (head == NULL || head->next == NULL) {
        return head;
    }

    while (head != NULL)
    {
        //临时存放的地址
        temp = head;
        //将 temp 从 head 中摘除,移位
        head = head->next;
        //将 temp 插入到 new_head 的头部
        temp->next = new_head;
        //临时temp存放到新节点
        new_head = temp;
    }
    return new_head;
}
```

## 4、就地逆置法反转链表

就地逆置法和头插法的实现思想类似，唯一的区别在于，头插法是通过建立一个新链表实现的，而就地逆置法则是直接对原链表做修改，从而实现将原链表反转。

值得一提的是，在原链表的基础上做修改，需要额外借助 2 个指针（假设分别为 beg 和 end）。仍以图 1 所示的链表为例，接下来用就地逆置法实现对该链表的反转：

![reverse_list8](../../../images/reverse_list8.PNG)

![reverse_list9](../../../images/reverse_list9.PNG)

由此，就实现了对图 1 链表的反转。 

具体实现代码如下：

```c++
link * local_reverse(link * head) {
    link * beg = NULL;
    link * end = NULL;
    if (head == NULL || head->next == NULL) {
        return head;
    }
    beg = head;
    end = head->next;
    while (end != NULL) {
        //将 end 从链表中摘除
        beg->next = end->next;
        //将 end 移动至链表头
        end->next = head;
        head = end;
        //调整 end 的指向，另其指向 beg 后的一个节点，为反转下一个节点做准备
        end = beg->next;
    }
    return head;
}
```

## 5、总结

本节仅以无头节点的链表为例，讲解了实现链表反转的 4 种方法。实际上，对于有头节点的链表反转：

使用迭代反转法实现时，初始状态忽略头节点（直接将 mid 指向首元节点），仅需在最后一步将头节点的 next 改为和 mid 同向即可；
使用头插法或者就地逆置法实现时，仅需将要插入的节点插入到头节点和首元节点之间即可；
递归法并不适用反转有头结点的链表（但并非不能实现），该方法更适用于反转无头结点的链表。
结合以上说明，读者可尝试修改本节代码，使它们能用于反转带头节点的链表。对于反转没有头节点的链表，读者可从反转无头节点链表下载；反之，对于采用迭代法、头插法以及就地逆置法反转有头节点的链表，读者可从反转有头节点链表处下载。