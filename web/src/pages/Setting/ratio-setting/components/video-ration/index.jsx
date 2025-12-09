import React from 'react';
import { Button } from '@douyinfe/semi-ui';
import { useWsPaginationTable } from '@hooks';
import { WsPaginationTable } from '@components';
import { editModal } from './components/edit';
import service from './service';
import { columns } from './columns';
import * as console from 'node:console';

export default function ModelSettingsVisualEditor(props) {
  const table = useWsPaginationTable(service.getWsVideoRationPageList);

  return (
    <div>
      <Button
        onClick={() => {
          editModal.open({}, () => {
            console.log('#');
          });
        }}
      >
        <>测试</>
      </Button>
      <WsPaginationTable {...table.tableProps} columns={columns()} />
    </div>
  );
}
