Vue.component("plan-user-selector", {
	props: ["v-user-id"],
	data: function () {
		return {}
	},
	methods: {
		change: function (userId) {
			this.$emit("change", userId)
		}
	},
	template: `<div>
	<user-selector :v-user-id="vUserId" data-url="/plans/users/options" @change="change"></user-selector>
</div>`
})