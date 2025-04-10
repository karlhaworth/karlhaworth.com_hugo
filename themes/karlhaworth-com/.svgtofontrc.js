import path from "path";
/**
 * @type {import('svgtofont').SvgToFontOptions}
 */
export default {
  fontName: "kh-font",
  css: {
    cssPath: "/kh-font/",
    hasTimestamp: false
  },
  generateInfoData: true,
  output: "assets/kh-font",
  sources: "static/svgs",
  emptyDist: true,
  outSVGReact: false
}