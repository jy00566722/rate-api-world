import { defineConfig } from 'wxt';
import vue from '@vitejs/plugin-vue';
import vueI18n from '@intlify/unplugin-vue-i18n/vite';

// See https://wxt.dev/api/config.html
export default defineConfig({
  imports: {
    addons: {
      vueTemplate: true,
    },
  },
  vite: () => ({
    plugins: [
      vue(),
      vueI18n({
        include: 'assets/locales/*.json',
      }),
    ]
  }),
  manifest: {
    name: '__MSG_extName__',
    version: '1.3.1',
    version_name: "1.3.1",
    description: '__MSG_description__',
    icons: {
      16: 'icons/16.png',
      48: 'icons/48.png',
      128: 'icons/128.png',
      256: 'icons/256.png',
    },
    default_locale: 'en_US',
    permissions: ["storage","alarms"],
    incognito:"spanning",
    host_permissions:[
      "https://rate-api.oeo.li/"
  ],
  //   optional_host_permissions:[
  //     // "https://*/*",
  //     // "http://*/*"
  // ],
    action: {
    default_title: "__MSG_actionTitle__",
    default_popup: "popup.html"
  },
    
  },
});
