import { createI18n } from 'vue-i18n';
// Use your default locale when importing the schema
import type schema from '~/assets/locales/en_US.json';
import messages from '@intlify/unplugin-vue-i18n/messages';

export type I18nSchema = typeof schema;
export type I18nLocales =  'en_US';

export default createI18n<[I18nSchema], I18nLocales>({
  messages: messages as any,
  silentTranslationWarn: true,
});