#include "linux/init.h"
#include "linux/module.h"
#include "linux/sched.h" //for_each_process
MODULE_LICENSE("BSD");
static int pid = -1;
module_param(pid, int, S_IRUGO);
static int killd_init(void)
{
    struct task_struct *p;
    printk(KERN_ALERT "killd: force D status process to death");
    printk(KERN_ALERT "killd: pid=%d", pid);
    for_each_process(p){
        if(p->pid == pid){
            printk("killd: found");
            set_task_state(p, TASK_STOPPED);
            printk(KERN_ALERT "killd: aha, make it stopped");
            return 0;
        }
    }
    printk("not found");
    return 0;
}
static void killd_exit(void)
{
    printk(KERN_ALERT "killd: bye");
}
module_init(killd_init);
module_exit(killd_exit);
