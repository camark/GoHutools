import { defineUserConfig } from "vuepress";
import { viteBundler } from "@vuepress/bundler-vite";
import { hopeTheme } from "vuepress-theme-hope";

export default defineUserConfig({
  lang: "zh-CN",
  title: "GoHutool",
  description: "Go 版本的 Hutool 工具库",

  base: "/",

  bundler: viteBundler(),

  theme: hopeTheme({
    logo: "/logo.svg",

    navbar: [
      "/",
      "/guide/",
      "/api/",
      "/examples/",
    ],

    sidebar: {
      "/guide/": "structure",
      "/api/": "structure",
      "/examples/": "structure",
    },

    locales: {
      "/": {
        navbar: [
          { text: "首页", link: "/" },
          { text: "指南", link: "/guide/" },
          { text: "API", link: "/api/" },
          { text: "示例", link: "/examples/" },
        ],
        sidebar: {
          "/guide/": [
            {
              text: "指南",
              children: [
                "/guide/",
                "/guide/getting-started.md",
                "/guide/install.md",
              ],
            },
          ],
          "/api/": [
            {
              text: "基础工具",
              children: [
                "/api/strutil.md",
                "/api/numutil.md",
                "/api/collutil.md",
                "/api/maputil.md",
                "/api/arrayutil.md",
                "/api/dateutil.md",
                "/api/convert.md",
                "/api/validate.md",
              ],
            },
            {
              text: "IO 与文件",
              children: [
                "/api/ioutil.md",
                "/api/fileutil.md",
                "/api/charsetutil.md",
              ],
            },
            {
              text: "网络与数据",
              children: [
                "/api/httpclient.md",
                "/api/jsonutil.md",
              ],
            },
            {
              text: "安全与编码",
              children: [
                "/api/crypto.md",
                "/api/codec.md",
                "/api/idutil.md",
                "/api/randomutil.md",
              ],
            },
            {
              text: "高级功能",
              children: [
                "/api/cache.md",
                "/api/cron.md",
                "/api/captcha.md",
                "/api/bloom.md",
                "/api/pool.md",
              ],
            },
            {
              text: "系统与配置",
              children: [
                "/api/log.md",
                "/api/setting.md",
                "/api/system.md",
                "/api/objectutil.md",
                "/api/regexutil.md",
              ],
            },
          ],
        },
      },
    },

    plugins: {
      mdEnhance: {
        demo: true,
        include: true,
        mark: true,
        chart: true,
        echarts: true,
        tabs: true,
        codetabs: true,
        tasklist: true,
      },
    },
  }),
});
