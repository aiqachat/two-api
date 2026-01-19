// 开发调试页面
import React from 'react';
import { Button, Toast } from '@douyinfe/semi-ui';
import { wsDev } from '@/helpers/ws-dev.js';

const DevPage = () => {
  return (
    <div style={{
      width: '100%',
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      marginTop: '30vh'
    }}>
      <Button
        type='primary'
        onClick={async () => {
          if(location.host !== 'localhost:5173') {
            return
          }
          await wsDev.autoLogin()
          Toast.success('已自动登录')
        }}
        size="large"
      >
        一键登录
      </Button>
    </div>
  );
};

export default DevPage;