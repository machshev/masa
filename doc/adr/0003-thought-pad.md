# 3. Thought Pad

Date: 2024-04-20

## Status

Accepted



## Context

There are many different todo apps out there. Some provide very simplistic
checklist style input with very little else in terms of meta information for
each todo. Others have many different fields that need adding to a todo item
before it can be created and have any sensible meaning in the system. The extra
fields such as priorities, start date, end date, Eisenhower Matrix (Urgancy,
Importance), reminder schedules, descriptions, notes, progress tracking and
other things. With GTD apps things like categories, where you need to be to
complete the task, what "goals" the task contributes towards.

Personally I find the more complicated systems tedious when it comes to quickly
jotting down a list of thoughts, and the ideas might not even make it to a full
todo task. These are the kind of things you might jot down on a pad of paper and
then from there convert to calander items or start a "Project" in a GTD todo app
or ticketing system in Github or Jira.

Clearly if the thought progresses to a full project with a complex hieracy of tasks
where some are time bound or dependent on events, then a list on a bit of paper
is not going to cut it. Now where did that bit of paper go?

The process that I tend to follow is something like the following. Although not
every thought warrants all the steps below, in fact I'd be quite worried about a
process that regularly had all these steps. Masa needs to support a workflow that
has some of the following for different types of use cases.

- Initial loose thoughts on paper, whiteboard, some unconstrain thought capture.
- Select ideas that are going to be acted on and discard others.
- Prioritise and catagorise ideas
- Refine the ideas:
    - Split up ideas into subtasks
    - Add detailed descriptions
    - Add periodic notes in the refinement process, maybe a meeting outcome?
    - Add links and other meta information
    - Add dependencies between tasks
- Track progress:
    - Add status (Proposal, Approved, Todo, Waiting, In Progress, Done)
    - Add event (Meeting, Phone call, ...)
    - Time tracking (Hours worked on task)
- Schedules and reminders
- Reportitng:
    - Tasks complete
    - Outstanding tasks
    - ...

Masa needs to scale from capturing adhoc thoughts while reading a book,
household chores, shopping lists, and things like that. All the way up to
colaborative projects with multi-person teams.

## Decision

Masa will provide a jot pad like temporary data store, which we are calling the
ThoughtPad. The ThoughtPad will be a simple short term idea capture system,
separate from the main database system required to support the more complicated
Masa workflows. It will store only a creation time stamp, and freeform text, no
other meta information. It's not intended to be interpreted by machines, but
only a staging ground or inbox for the main Masa database system.

The ThoughtPad will support a triage workflow, making it easy to modify,
combine, and delete ideas. This is not longterm storage, but more of an inbox
for the rest of the Masa system. It's expected that the items in this list will
be regularly cleared out by either promoting them to the main database where
they can be further refined, or deleting them all together.

Users will be encouraged to process the ThoughtPad items regularly and not leave
items on there for long periods of time. This would likely result in stale data
and notes left in short form that may not be understandable in a few days time.

### Data storage

There will only be one ThoughtPad per device, where a user has multiple devices
at some point we need to work out how the data would be synchronised.

The ThoughtPad will be an in memory data store to allow for simpler operations.
This will be peristed to the file system in the XDG_DATADIR when the masa server
is shut down. Each operation on the ThoughtPad will be appended to a transaction
log file while the server is running. Which will allow recovery of state if the
server crashes.

On server shutdown, any existing ThoughtPad file will be renamed backup. Then
the current in memory ThoughtPad will be persisted, and both the transaction log
file and the backup removed.

When the server starts it will check to see if there are any existing ThoughtPad
files:

1. Load the ThoughtPad data from disk:
    - If no files it will create a new empty in memory ThoughtPad.
    - If a backup ThoughtPad exists then:
        - Read it into memory
        - If a transaction file exists then apply the transactions
    - If only a main ThoughtPad file exists then load it into memory
2. Save the new in memory ThoughtPad to the main ThoughtPad file
3. Remove the backup ThoughtPad file and the old transaction file
4. Open a new transaction file for the current session

This way if the server crashes for any reason (looses power) then the data can
be recovered reliably.

## Consequences

Seporating the ThoughtPad system from the main Masa system simplifies the design
significantly. It avoids the trap of trying to reuse the complicated data
structures for the simple use case.

Using an in memory store backed by transaction file should mean thoughts are
reliably captured and not lost.

