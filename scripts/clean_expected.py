#!/usr/bin/env python3
"""Clean expected YAML files by removing dynamic fields that change between runs."""

import sys
import yaml


def clean_step(step):
    """Clean a step, removing dynamic fields and recursively cleaning nested steps."""
    cleaned = {}

    if "data" in step:
        data = {}
        if step["data"].get("action"):
            data["action"] = step["data"]["action"]
        if step["data"].get("expected_result"):
            data["expected_result"] = step["data"]["expected_result"]
        if step["data"].get("input_data"):
            data["input_data"] = step["data"]["input_data"]
        if data:
            cleaned["data"] = data

    if "execution" in step and step["execution"].get("status"):
        cleaned["execution"] = {"status": step["execution"]["status"]}

    if "steps" in step and step["steps"]:
        cleaned["steps"] = [clean_step(s) for s in step["steps"]]

    return cleaned


def clean_attachment(att):
    """Clean an attachment, keeping only file_name and mime_type."""
    cleaned = {}
    if att.get("file_name"):
        cleaned["file_name"] = att["file_name"]
    if att.get("mime_type"):
        cleaned["mime_type"] = att["mime_type"]
    return cleaned


def clean_result(result):
    """Clean a result, removing dynamic fields."""
    cleaned = {}

    if result.get("title"):
        cleaned["title"] = result["title"]

    if result.get("signature"):
        cleaned["signature"] = result["signature"]

    if result.get("testops_ids"):
        cleaned["testops_ids"] = result["testops_ids"]

    # Promote status from execution to top level
    status = None
    if result.get("status"):
        status = result["status"]
    elif result.get("execution", {}).get("status"):
        status = result["execution"]["status"]
    if status:
        cleaned["status"] = status

    # Keep non-empty fields
    if result.get("fields") and result["fields"] != {}:
        cleaned["fields"] = result["fields"]

    # Keep non-empty params
    if result.get("params") and result["params"] != {}:
        cleaned["params"] = result["params"]

    # Keep relations
    if result.get("relations"):
        cleaned["relations"] = result["relations"]

    # Clean steps (remove empty)
    if result.get("steps") and result["steps"] != []:
        cleaned["steps"] = [clean_step(s) for s in result["steps"]]

    # Clean attachments (remove empty, keep only file_name and mime_type)
    if result.get("attachments") and result["attachments"] != []:
        cleaned["attachments"] = [clean_attachment(a) for a in result["attachments"]]

    # Keep muted only if true
    if result.get("muted"):
        cleaned["muted"] = True

    # Skip: message, stacktrace, execution block, param_groups (if null/empty),
    # id, start_time, end_time, duration, thread

    return cleaned


def clean_expected(data):
    """Clean the entire expected YAML data."""
    cleaned = {}

    if "run" in data:
        cleaned["run"] = data["run"]

    if "results" in data:
        cleaned["results"] = [clean_result(r) for r in data["results"]]

    return cleaned


def main():
    if len(sys.argv) < 2:
        print("Usage: clean_expected.py <file1.yaml> [file2.yaml ...]")
        sys.exit(1)

    for filepath in sys.argv[1:]:
        with open(filepath, "r") as f:
            data = yaml.safe_load(f)

        cleaned = clean_expected(data)

        with open(filepath, "w") as f:
            yaml.dump(
                cleaned,
                f,
                default_flow_style=False,
                allow_unicode=True,
                sort_keys=False,
            )

        print(f"Cleaned: {filepath}")


if __name__ == "__main__":
    main()
