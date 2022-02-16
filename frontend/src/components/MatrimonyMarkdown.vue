<template>
  <vue-markdown
    :source="source"
    class="prose prose-slate dark:prose-invert mx-auto"
  />
</template>

<script setup>
import VueMarkdown from "vue-markdown-render";

defineProps({
  source: { type: String, default: "" },
});
</script>

<script>
export default {
  name: "MatrimonyMarkdown",
};
</script>
