import React from 'react';
import { Button, Form } from '@douyinfe/semi-ui';
import { useWsPaginationTable } from '@hooks';
import { WsPaginationSearch, WsPaginationTable } from '@components';
import { editModal } from './components/edit';
import service from './service';
import { columns } from './columns';
import { IconPlus, IconSearch } from '@douyinfe/semi-icons';

export default function ModelSettingsVisualEditor() {
  const table = useWsPaginationTable(service.getWsVideoRationPageList);

  return (
    <div>
      <WsPaginationSearch {...table.searchProps}>
        <Button
          icon={<IconPlus />}
          size='small'
          onClick={() => {
            editModal.open({}, () => {
              table.reload();
            });
          }}
        >
          <>添加模型</>
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
      <WsPaginationTable
        {...table.tableProps}
        columns={columns(() => table.refresh())}
      />
    </div>
  );
}
