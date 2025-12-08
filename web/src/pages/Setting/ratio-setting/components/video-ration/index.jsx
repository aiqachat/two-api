import React from 'react';
import {
  Button,
} from '@douyinfe/semi-ui';
import service from './service'

export default function ModelSettingsVisualEditor(props) {
  console.log(props)
  return (
    <div>
      <Button
        onClick={() => {
          service.createWsVideoRation({}).then()
        }}
      >
        <>测试</>
      </Button>
    </div>
  );
}
