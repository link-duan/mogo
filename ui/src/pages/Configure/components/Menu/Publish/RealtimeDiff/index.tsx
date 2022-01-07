import { Button, Spin } from "antd";
import { useEffect } from "react";
import { MonacoDiffEditor } from "react-monaco-editor";
import { useModel } from "@@/plugin-model/useModel";
import CustomModal from "@/components/CustomModal";
import diffStyles from "@/pages/Configure/components/Menu/Publish/RealtimeDiff/index.less";

type RealtimeDiffProps = {
  visible?: boolean;
  configId?: number;
  version?: string;
  onCancel?: () => void;
  onOk: () => void;
};

const RealtimeDiff = (props: RealtimeDiffProps) => {
  const { visible, configId, version, onCancel, onOk } = props;
  const {
    selectedClusterId,
    selectedConfigMap,
    selectedNameSpace,
    configurationList,
    doGetCurrentVersionConfiguration,
    doGetOnlineConfiguration,
    doPublishConfiguration,
  } = useModel("configure");

  useEffect(() => {
    if (!visible || !configId || !version) return;
    doGetCurrentVersionConfiguration.run(configId, version);
    const config = configurationList.find((item) => item.id === configId);
    doGetOnlineConfiguration.run(
      selectedClusterId as number,
      selectedNameSpace as string,
      selectedConfigMap as string,
      `${config?.name}.${config?.format}`
    );
  }, [visible, configId, version]);

  return (
    <CustomModal
      visible={visible}
      title="实时配置 Diff"
      width="90%"
      footer={
        <Button
          loading={doPublishConfiguration.loading}
          size={"small"}
          type={"primary"}
          onClick={() => onOk()}
        >
          发布
        </Button>
      }
      onCancel={onCancel}
    >
      <Spin
        spinning={
          doGetCurrentVersionConfiguration.loading ||
          doGetOnlineConfiguration.loading
        }
      >
        <div className={diffStyles.diffHeader}>
          <div className={diffStyles.title}>生效中配置</div>
          <div className={diffStyles.title}>本次发布配置</div>
        </div>
        <MonacoDiffEditor
          language="json"
          theme="vs-dark"
          original={doGetOnlineConfiguration.data}
          value={doGetCurrentVersionConfiguration.data?.content}
          height="70vh"
          options={{
            automaticLayout: true,
            scrollBeyondLastLine: false,
          }}
        />
      </Spin>
    </CustomModal>
  );
};

export default RealtimeDiff;
