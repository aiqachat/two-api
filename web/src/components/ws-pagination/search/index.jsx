import { Button, Card, Form, Space } from '@douyinfe/semi-ui';
import { useRef } from 'react';

export const WsPaginationSearch = ({ onSearch, onReset, children }) => {
  const formRef = useRef(null);

  return (
    <Card>
      <Form ref={formRef}>
        <Space>
          {children}
          <Button
            size='small'
            type='primary'
            onClick={async () => {
              const values = await formRef.current?.formApi.validate();
              onSearch(values);
            }}
          >
            搜索
          </Button>
          <Button
            size='small'
            onClick={() => {
              formRef.current?.formApi.reset();
              onReset();
            }}
          >
            重置
          </Button>
        </Space>
      </Form>
    </Card>
  );
};
