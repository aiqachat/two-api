// 开发调试页面
import React from 'react';
import { Button, Space, Toast } from '@douyinfe/semi-ui';
import { WsError, wsDev } from '@/helpers';
import service from './service';

const DevPage = () => {
  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        marginTop: '30vh',
      }}
    >
      <Space>
        {location.host === 'localhost:5173' && (
          <Button
            type='primary'
            onClick={async () => {
              await wsDev.autoLogin();
              Toast.success('已自动登录');
            }}
            size='large'
          >
            一键登录
          </Button>
        )}
        <Button
          type='primary'
          onClick={async () => {
            try {
              await service.fixVideoRatioConfig();
              Toast.success('视频倍率配置已修复成功');
            } catch (e) {
              WsError.handleError(e);
            }
          }}
          size='large'
        >
          修复视频倍率配置
        </Button>
      </Space>
    </div>
  );
};

export default DevPage;
