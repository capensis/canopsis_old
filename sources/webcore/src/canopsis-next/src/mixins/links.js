import { isObject } from 'lodash';

export default {
  methods: {
    /*
    * The linkbuilders used to return the links directly as
    * strings. They can now also return objects with the
    * properties 'label' and 'link', allowing to change the link's
    * label.
    * The following code converts the "legacy" representation
    * (strings) into the "new" representation, so they can be
    * displayed in the same manner by the template.
    */
    harmonizeLinks(links, category) {
      return links.map((link, index) => {
        if (isObject(link) && link.link && link.label) {
          return link;
        }

        return {
          label: `${category} - ${index}`,
          link,
        };
      });
    },
  },
};
