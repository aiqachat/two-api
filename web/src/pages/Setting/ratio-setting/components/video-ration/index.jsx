import React from 'react';
import { Button, Form } from '@douyinfe/semi-ui';
import { useWsPaginationTable } from '@hooks';
import { WsPaginationSearch, WsPaginationTable } from '@components';
import { editModal } from './components/edit';
import service from './service';
import { columns } from './columns';
import { IconSearch } from '@douyinfe/semi-icons';

export default function ModelSettingsVisualEditor(props) {
  const table = useWsPaginationTable(service.getWsVideoRationPageList);

  return (
    <div>
      <WsPaginationSearch {...table.searchProps}>
        <Button
          size='small'
          onClick={() => {
            editModal.open({}, () => {
              console.log('#');
            });
          }}
        >
          <>测试</>
        </Button>
        <Form.Input
          field='model_name'
          prefix={<IconSearch />}
          placeholder='搜索模型名称'
          showClear
          pure
          size='small'
        />
      </WsPaginationSearch>
      <WsPaginationTable {...table.tableProps} columns={columns()} />
    </div>
  );
}
