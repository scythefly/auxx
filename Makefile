main:
	bash build.sh

leetcode:
	bash plugins/leetcode/build.sh ${CURDIR}/bin/leetcode.so

utility:
	bash plugins/utility/build.sh ${CURDIR}/bin/utility.so

ui:
	bash plugins/ui/build.sh ${CURDIR}/bin/ui.so