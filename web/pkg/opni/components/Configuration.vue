<script>
import jsyaml from 'js-yaml';
import YamlEditor from '@shell/components/YamlEditor';
import Loading from '@shell/components/Loading';
import AsyncButton from '@shell/components/AsyncButton';
import { Banner } from '@components/Banner';
import { GatewayConfigSpec } from '@pkg/opni/generated/github.com/rancher/opni/pkg/config/v1/gateway_config_pb';
import { exceptionToErrorsArray } from '../utils/error';
import { getGatewayConfig, updateGatewayConfig } from '../utils/requests/management';
import { Config, DriverUtil } from '../api/opni';

// replaces BigInts with Numbers
function sanitizeBigInts(obj) {
  if (typeof obj === 'bigint') {
    return Number(obj);
  }
  if (Array.isArray(obj)) {
    return obj.map(sanitizeBigInts);
  }
  if (obj && typeof obj === 'object') {
    return Object.fromEntries(Object.entries(obj).map(([k, v]) => [k, sanitizeBigInts(v)]));
  }

  return obj;
}
export default {
  components: {
    YamlEditor,
    AsyncButton,
    Loading,
    Banner
  },

  async fetch() {
    await this.load();
  },

  data() {
    return {
      loading:        false,
      config:         null,
      editorContents: '',
      error:          '',
    };
  },

  methods: {
    async load() {
      try {
        this.loading = true;
        this.$set(this, 'config', await Config.service.GetConfiguration(new DriverUtil.types.GetRequest()));
        const contents = jsyaml.dump(sanitizeBigInts(structuredClone(this.config)));

        this.$set(this, 'editorContents', contents);
      } catch (err) {
        this.$set(this, 'error', exceptionToErrorsArray(err).join('; '));
      } finally {
        this.loading = false;
      }
    },
    async save(buttonCallback) {
      try {
        const obj = jsyaml.load(this.editorContents);

        await Config.service.SetConfiguration(new Config.types.SetRequest({ spec: new GatewayConfigSpec(obj) }));
        await this.load();
      } catch (err) {
        buttonCallback(false);
        this.$set(this, 'error', exceptionToErrorsArray(err).join('; '));
      } finally {
        buttonCallback(!this.error);
      }
    },
    async reset() {
      this.$set(this, 'error', '');
      await this.load();
    },
  }
};
</script>
<template>
  <Loading v-if="loading || $fetchState.pending" />
  <div v-else>
    <header class="m-0">
      <div class="title">
        <h1>Configuration</h1>
      </div>
    </header>
    <Banner
      v-if="error"
      color="error"
      :label="error"
    />
    <div slot="body">
      <YamlEditor
        v-model="editorContents"
        :as-object="false"
        :editor-mode="'EDIT_CODE'"
        :read-only="false"
        class="yaml-editor"
      />
      <div class="row actions-container mt-10">
        <button class="btn role-secondary mr-10" @click="reset">
          Reset
        </button>
        <AsyncButton
          class="btn role-primary mr-10"
          action-label="Save"
          waiting-label="Saving..."
          success-label="Success"
          error-label="Error"
          @click="save"
        />
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.actions-container {
  display: flex;
  justify-content: flex-end;
}
.buttons {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  width: 100%;
}
.error-message {
  color: var(--error);
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
