import React from 'react';
import { Button } from '@douyinfe/semi-ui';
import { editModal } from './components/edit';

export default function ModelSettingsVisualEditor(props) {
  console.log(props);
  return (
    <div>
      <Button
        onClick={() => {
          editModal.open({}, () => {
            console.log('#')
          });
        }}
      >
        <>测试</>
      </Button>
    </div>
  );
}
