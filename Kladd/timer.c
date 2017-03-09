#include <time.h>
#include <assert.h>

static time_t g_startTime = -1;

void
timer_start(){
    g_startTime = time(0);
    assert(g_startTime != -1);
}


void 
timer_stop(){
  g_startTime = -1;
}

int 
timer_isTimeOut(){
  if(g_startTime <0){
    // There is no timeout, because the timer is not started
    return 0;
  }

  time_t now = time(0);
  if(now - g_startTime > 3){
    return 1;
  }else{
    return 0;
  }
}
