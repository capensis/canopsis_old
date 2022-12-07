import { QUICK_RANGES } from '@/constants';

export default {
  title: 'Quick ranges',
  timeField: 'Time field',
  types: {
    [QUICK_RANGES.custom.value]: 'Custom',
    [QUICK_RANGES.last15Minutes.value]: 'Last 15 minutes',
    [QUICK_RANGES.last30Minutes.value]: 'Last 30 minutes',
    [QUICK_RANGES.last1Hour.value]: 'Last 1 hour',
    [QUICK_RANGES.last3Hour.value]: 'Last 3 hour',
    [QUICK_RANGES.last6Hour.value]: 'Last 6 hour',
    [QUICK_RANGES.last12Hour.value]: 'Last 12 hour',
    [QUICK_RANGES.last24Hour.value]: 'Last 24 hour',
    [QUICK_RANGES.last2Days.value]: 'Last 2 days',
    [QUICK_RANGES.last7Days.value]: 'Last 7 days',
    [QUICK_RANGES.last30Days.value]: 'Last 30 days',
    [QUICK_RANGES.last1Year.value]: 'Last 1 year',
    [QUICK_RANGES.yesterday.value]: 'Yesterday',
    [QUICK_RANGES.previousWeek.value]: 'Previous week',
    [QUICK_RANGES.previousMonth.value]: 'Previous month',
    [QUICK_RANGES.today.value]: 'Today',
    [QUICK_RANGES.todaySoFar.value]: 'Today so far',
    [QUICK_RANGES.thisWeek.value]: 'This week',
    [QUICK_RANGES.thisWeekSoFar.value]: 'This week so far',
    [QUICK_RANGES.thisMonth.value]: 'This month',
    [QUICK_RANGES.thisMonthSoFar.value]: 'This month so far',
  },
};
