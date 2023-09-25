import linkifyHtmlLib from 'linkify-html';
import sanitizeHtmlLib from 'sanitize-html';

const DEFAULT_SANITIZE_OPTIONS = {
  allowedTags: [
    'h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'blockquote', 'p', 'a', 'ul', 'ol', 'nl', 'li', 'b', 'i', 'strong', 'em',
    'code', 'hr', 'br', 'div', 'table', 'thead', 'caption', 'tbody', 'tr', 'th', 'td', 'pre', 'span', 'font', 'u',
    'img', 'video', 'audio', 'abbr', 'address', 'bdo', 'cite', 'q', 'dfn', 'var', 'area', 'map', 'dl', 'dt', 'dd',
    'section', 'article', 'colgroup', 'col', 'strike',

    /**
     * VUE COMPONENTS
     */
    'router-link', 'c-alarm-chip', 'c-alarm-tags-chips', 'c-copy-wrapper', 'c-links-list', 'service-entities-list',
  ],
  allowedAttributes: {
    '*': [
      'style', 'title', 'class', 'id', 'v-if', 'autoplay', 'colspan', 'controls', 'dir', 'align', 'width', 'height',
      'role',
    ],
    a: ['href', 'name', 'target'],
    img: ['src', 'alt'],
    font: ['color', 'size', 'face'],
    'router-link': ['href', 'name', 'target', 'to'],
    'c-alarm-chip': ['value'],
    'c-alarm-tags-chips': [':alarm', 'inline-count', '@select'],
    'c-copy-wrapper': ['value'],
    'c-links-list': [':links', ':category'],
    'service-entities-list': [
      ':service', ':service-entities', ':widget-parameters', ':pagination', ':total-items', 'entity-name-field',
      '@refresh', '@update:pagination',
    ],
  },
};

const DEFAULT_LINKIFY_OPTIONS = { target: '_blank' };

/**
 * Sanitize HTML document
 *
 * @param {string} [html = '']
 * @param {Object} [options = DEFAULT_SANITIZE_OPTIONS]
 * @return {string}
 */
export const sanitizeHtml = (html = '', options = DEFAULT_SANITIZE_OPTIONS) => sanitizeHtmlLib(html, options);

/**
 * Convert all links in html to tag <a>
 *
 * @param {string} [html = '']
 * @param {Object} [options = DEFAULT_LINKIFY_OPTIONS]
 * @return {string}
 */
export const linkifyHtml = (html = '', options = DEFAULT_LINKIFY_OPTIONS) => linkifyHtmlLib(html, options);
