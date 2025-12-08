import { Toast } from '@douyinfe/semi-ui';
import { debounce } from 'lodash';
import { isAxiosError } from 'axios';

export const openErrorDialog = debounce(
  (msg) => {
    Toast.error(msg);
  },
  200,
  {
    leading: true,
    trailing: false,
  },
);

export class WsError extends Error {
  constructor(message, code = -1) {
    super(message);
    this.message = message;
    this.code = code;
    this.isServiceError = true;
  }

  /**
   * 抛出业务异常
   * @param message 错误消息
   * @param code 错误号码
   */
  static throw(message, code) {
    throw new WsError(message, code);
  }

  /**
   * 抛出业务异常
   * @param expression 表达式
   * @param message 错误消息
   * @param code 错误号码
   */
  static throwIf(expression, message, code) {
    if (!expression) return;
    throw new WsError(message, code);
  }

  // 解决业务异常
  static handleError(e) {
    // axios的异常已被API模块处理
    if (isAxiosError(e)) {
      return;
    }
    console.error(e);
    if (e?.isServiceError) {
      openErrorDialog('错误：' + e.message);
      return;
    }
    openErrorDialog('错误：' + e);
  }

  // 检查Api结果
  static checkApiResult(res) {
    if (res.success !== true) {
      WsError.throw(res.message);
    }
  }
}
